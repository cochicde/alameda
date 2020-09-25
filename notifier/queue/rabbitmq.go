package queue

import (
	"encoding/json"
	"os"
	"time"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"prophetstor.com/alameda/notifier/notifying"
	notifier_utils "prophetstor.com/alameda/notifier/utils"
	datahubpkg "prophetstor.com/alameda/pkg/datahub"
	k8s_utils "prophetstor.com/alameda/pkg/utils/kubernetes"
	"prophetstor.com/alameda/pkg/utils/log"
	datahub_events "prophetstor.com/api/datahub/events"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var scope = log.RegisterScope("queue", "queue", 0)

type rabbitmqClient struct {
	conn          *amqp.Connection
	datahubClient *datahubpkg.Client
	mgr           manager.Manager
}

func NewRabbitMQClient(mgr manager.Manager, queueURL string, datahubClient *datahubpkg.Client) *rabbitmqClient {
	retryConn := viper.GetInt("rabbitmq.connRetry")
	retryConnInterval := viper.GetInt64("rabbitmq.connRetryInterval")
	for retry := 0; retry < retryConn; retry++ {
		conn, err := amqp.Dial(queueURL)
		if err == nil {
			return &rabbitmqClient{
				conn:          conn,
				datahubClient: datahubClient,
				mgr:           mgr,
			}
		}
		if err != nil {
			if retry == retryConn-1 {
				scope.Errorf("failed to connect to queue %s: %s", queueURL, err.Error())
				os.Exit(1)
			} else {
				scope.Errorf("failed to connect to queue %s: %s. Retry after %d seconds",
					queueURL, err.Error(), retryConnInterval)
				time.Sleep(time.Duration(retryConnInterval) * time.Second)
			}
		}

	}

	return &rabbitmqClient{
		datahubClient: datahubClient,
		mgr:           mgr,
	}
}

func (client *rabbitmqClient) Start() {
	ch, err := client.conn.Channel()
	defer ch.Close()
	if err != nil {
		scope.Errorf(err.Error())
	}

	q, err := ch.QueueDeclare(
		"event", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		scope.Errorf(err.Error())
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		scope.Errorf(err.Error())
	}
	forever := make(chan bool)
	scope.Infof("get cluster info - begin")
	clusterInfo, err := notifier_utils.GetClusterInfo(client.mgr.GetClient())
	if err != nil {
		scope.Errorf("unable to get cluster info: %s", err.Error())
	}
	scope.Infof("get cluster info - done")
	uid, err := k8s_utils.GetClusterUID(client.mgr.GetClient())
	if err != nil {
		scope.Errorf("unable to get cluster id: %s", err.Error())
	} else {
		clusterInfo.UID = uid
	}
	scope.Infof("clusterInfo: %#v", clusterInfo)
	notifier := notifying.NewNotifier(client.mgr, client.datahubClient, &clusterInfo)
	go func() {
		for d := range msgs {
			scope.Infof("Received events: %s", d.Body)
			evts := &[]*datahub_events.Event{}
			decodeErr := json.Unmarshal(d.Body, evts)
			if decodeErr != nil {
				scope.Error(decodeErr.Error())
			} else {
				notifier.NotifyEvents(*evts)
			}
			d.Ack(false)
		}
	}()
	<-forever
}
