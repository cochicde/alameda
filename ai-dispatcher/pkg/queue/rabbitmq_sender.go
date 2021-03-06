package queue

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"

	"github.com/containers-ai/alameda/ai-dispatcher/pkg/config"
	utils "github.com/containers-ai/alameda/ai-dispatcher/pkg/utils"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type retry struct {
	publishRetryTime       int
	publishRetryIntervalMS int64
	consumeRetryTime       int
	consumeRetryIntervalMS int64
}

type RabbitMQSender struct {
	conn       *amqp.Connection
	queueURL   string
	retryItvMS int64
	connNotify chan *amqp.Error
}

func (sender *RabbitMQSender) SendJsonString(queueName, jsonStr, msgID string, granularity int64) error {

	publishRetryIntervalMS := sender.getRetry().publishRetryIntervalMS
	var err error
	for start := time.Now(); time.Since(start) < (time.Duration(granularity) * time.Second); {

		err = sender.sendJob(queueName, jsonStr, msgID)
		if err == nil {
			scope.Infof("[%s] Send job successfully.", msgID)
			return nil
		} else {
			scope.Errorf("[%s] Send job failed due to %s. Retry job sending later if sending process does not reach timeout",
				msgID, err.Error())
		}
		time.Sleep(time.Duration(publishRetryIntervalMS) * time.Millisecond)
	}
	if err != nil {
		return err
	}
	return fmt.Errorf("unknown error to send message to queue %s", queueName)
}

func (sender *RabbitMQSender) sendJob(queueName, jsonStr, msgID string) error {
	if sender.conn.IsClosed() {
		scope.Warnf("try to send job but the connection is close, establish a new connection")
		sender.conn = GetQueueConn(sender.queueURL, sender.retryItvMS)
	}
	queueCH, err := sender.conn.Channel()
	if err != nil {
		return err
	}
	defer queueCH.Close()

	notifyClose := make(chan *amqp.Error)
	notifyConfirm := make(chan amqp.Confirmation)
	queueCH.Confirm(false)
	queueCH.NotifyClose(notifyClose)
	queueCH.NotifyPublish(notifyConfirm)

	q, err := queueCH.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp.Table{
			"x-message-deduplication": true,
		}, // arguments
	)
	if err != nil {
		return err
	}

	err = queueCH.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(jsonStr),
			DeliveryMode: 2, // 2 means persistent
			Headers: amqp.Table{
				//"x-deduplication-header": sender.getMessageHash(msgID),
				"x-deduplication-header": msgID,
			},
		})

	if err != nil {
		return err
	}

	ackTimeoutSec := viper.GetInt("queue.retry.ackTimeoutSec")
	if ackTimeoutSec <= 0 {
		ackTimeoutSec = DEFAULT_ACK_TIMEOUT_SEC
	}
	ticker := time.NewTicker(time.Duration(ackTimeoutSec) * time.Second)

	for {
		select {
		case confirm := <-notifyConfirm:
			if confirm.Ack {
				return nil
			} else {
				return fmt.Errorf("send job to queue %s failed, server does not receive the job", queueName)
			}
		case confirm := <-notifyClose:
			if confirm != nil {
				return fmt.Errorf(confirm.Error())
			}
			return fmt.Errorf("send jobs to queue %s with unknown exception of notification close", queueName)
		case confirm := <-sender.connNotify:
			if confirm != nil {
				if !sender.conn.IsClosed() {
					err := sender.conn.Close()
					if err != nil {
						scope.Errorf(err.Error())
					}
				}
				sender.conn = GetQueueConn(sender.queueURL, sender.retryItvMS)
				return fmt.Errorf(confirm.Error())
			} else {
				return nil
			}
		case <-ticker.C:
			err := queueCH.Close()
			if err != nil {
				return fmt.Errorf("send job to queue %s timeout and close channel failed: %s",
					queueName, err.Error())
			}

			return fmt.Errorf("send job to queue %s timeout", queueName)
		}
	}
}

func NewRabbitMQSender(queueURL string, retryItvMS int64) (*RabbitMQSender, *amqp.Connection) {
	conn := GetQueueConn(queueURL, retryItvMS)
	sender := &RabbitMQSender{
		queueURL:   queueURL,
		retryItvMS: retryItvMS,
		conn:       conn,
	}
	sender.connNotify = sender.conn.NotifyClose(make(chan *amqp.Error))
	return sender, sender.conn
}

func (sender *RabbitMQSender) getRetry() *retry {
	publishRetryTime := viper.GetInt("queue.retry.publishTime")
	if publishRetryTime == 0 {
		publishRetryTime = DEFAULT_PUBLISH_RETRY_TIME
	}

	publishRetryIntervalMS := viper.GetInt64("queue.retry.publishIntervalMs")
	if publishRetryIntervalMS == 0 {
		publishRetryIntervalMS = DEFAULT_PUBLISH_RETRY_INTERVAL_MS
	}
	return &retry{
		publishRetryTime:       publishRetryTime,
		publishRetryIntervalMS: publishRetryIntervalMS,
	}
}

func (sender *RabbitMQSender) getMessageHash(msgStr string) string {
	h := sha1.New()
	h.Write([]byte(msgStr))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (sender *RabbitMQSender) SendJob(queueName string, unit *config.Unit,
	columns, values []string, metricType datahub_common.MetricType, granularity int64) error {
	jobID, err := utils.GetJobID(unit, values, columns,
		metricType, granularity)
	if err != nil {
		return err
	}
	body := map[string]interface{}{
		"predictor":        unit.Predictor,
		"scopeStr":         fmt.Sprintf("%s", unit.Scope),
		"measurement":      unit.Measurement,
		"category":         unit.Category,
		"unit_type":        unit.Type,
		"granularity":      utils.GetGranularityStr(granularity),
		"granularitySec":   fmt.Sprintf("%v", granularity),
		"createTimestamp":  time.Now().Unix(),
		"metricType":       fmt.Sprintf("%d", metricType),
		"metricTypeString": fmt.Sprintf("%s", metricType),
		"jobID":            jobID,
		"object_meta":      utils.GetJobMap(values, columns, metricType, granularity),
	}

	bodyStr, err := json.Marshal(body)
	if err != nil {
		return err
	}
	scope.Debugf(fmt.Sprintf("queue %s: %s", queueName, string(bodyStr)))
	err = sender.SendJsonString(queueName, string(bodyStr), jobID, granularity)
	return err
}
