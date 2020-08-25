package prometheus

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

const (
	EndpointPrefix       = "/api/v1"
	ExpressionQuery      = "/query"
	ExpressionQueryRange = "/query_range"
)

// Query data over a range of time from prometheus
func (p *Prometheus) Query(ctx context.Context, query string, startTime, timeout *time.Time) (Response, error) {

	var (
		err error

		endpoint        = EndpointPrefix + ExpressionQuery
		queryParameters = url.Values{}

		u            *url.URL
		httpRequest  *http.Request
		httpResponse *http.Response

		response Response
	)

	queryParameters.Set("query", query)

	if startTime != nil {
		queryParameters.Set("time", strconv.FormatInt(int64(startTime.Unix()), 10))
	}

	if timeout != nil {
		queryParameters.Set("timeout", strconv.FormatInt(int64(timeout.Unix()), 10))
	}

	u, err = url.Parse(p.Config.URL + endpoint)
	if err != nil {
		scope.Errorf("failed to parse url: %s", err.Error())
		scope.Error("failed to query prometheus")
		return Response{}, errors.New("failed to parse prometheus url")
	}
	u.RawQuery = queryParameters.Encode()

	httpRequest, err = http.NewRequest("GET", u.String(), nil)
	if err != nil {
		scope.Errorf("failed to create http request: %s", err.Error())
		scope.Error("failed to query prometheus")
		return Response{}, errors.New("failed to create http request")
	}
	if token := p.Config.bearerToken; token != "" {
		h := http.Header{
			"Authorization": []string{fmt.Sprintf(" Bearer %s", token)},
		}
		httpRequest.Header = h
	}

	httpRequest = httpRequest.WithContext(ctx)
	httpResponse, err = p.Client.Do(httpRequest)
	if err != nil {
		scope.Errorf("failed to send http request to prometheus: %s", err.Error())
		scope.Error("failed to query prometheus")
		return Response{}, errors.New("failed to send http request to prometheus")
	}
	err = decodeHTTPResponse(httpResponse, &response)
	if err != nil {
		return Response{}, errors.Wrap(err, "failed to query prometheus")
	}

	defer p.Close()

	return response, nil
}

// QueryRange data over a range of time from prometheus
func (p *Prometheus) QueryRange(ctx context.Context, query string, startTime, endTime *time.Time, stepTime *time.Duration) (Response, error) {

	var (
		err error

		startTimeString string
		endTimeString   string
		stepTimeString  string

		endpoint        = EndpointPrefix + ExpressionQueryRange
		queryParameters = url.Values{}

		u            *url.URL
		httpRequest  *http.Request
		httpResponse *http.Response

		response Response
	)

	if startTime == nil {
		tmpTime := time.Unix(0, 0)
		startTime = &tmpTime
	}
	startTimeString = strconv.FormatInt(int64(startTime.Unix()), 10)

	if endTime == nil {
		tmpTime := time.Now()
		endTime = &tmpTime
	}
	endTimeString = strconv.FormatInt(int64(endTime.Unix()), 10)

	if stepTime == nil {
		copyDefaultStepTime := DBCommon.DefaultStepTime
		stepTime = &copyDefaultStepTime
	}
	stepTimeString = strconv.FormatInt(int64(stepTime.Nanoseconds()/int64(time.Second)), 10)

	queryParameters.Set("query", query)
	queryParameters.Set("start", startTimeString)
	queryParameters.Set("end", endTimeString)
	queryParameters.Set("step", stepTimeString)

	u, err = url.Parse(p.Config.URL + endpoint)
	if err != nil {
		scope.Errorf("failed to parse prometheus url: %s", err.Error())
		scope.Error("failed to query_range prometheus")
		return Response{}, errors.New("failed to parse prometheus url")
	}
	u.RawQuery = queryParameters.Encode()

	httpRequest, err = http.NewRequest("GET", u.String(), nil)
	if err != nil {
		scope.Errorf("failed to create http request: %s", err.Error())
		scope.Error("failed to query_range prometheus")
		return Response{}, errors.New("failed to create http request")
	}
	if token := p.Config.bearerToken; token != "" {
		h := http.Header{
			"Authorization": []string{fmt.Sprintf(" Bearer %s", strings.TrimSpace(token))},
		}
		httpRequest.Header = h
	}

	responseChan := make(chan *http.Response)
	wg, wgCTX := errgroup.WithContext(ctx)
	httpRequest = httpRequest.WithContext(wgCTX)
	wg.Go(func() error {
		httpResponse, err = p.Client.Do(httpRequest)
		if err != nil {
			scope.Errorf("failed to send http request to prometheus: %s", err.Error())
			scope.Error("failed to query_range prometheus")
			return errors.Wrap(err, "failed to send http request to prometheus")
		}
		responseChan <- httpResponse
		return nil
	})

	select {
	case <-wgCTX.Done():
		return Response{}, wg.Wait()
	case httpResponse = <-responseChan:
	}

	err = decodeHTTPResponse(httpResponse, &response)
	if err != nil {
		return Response{}, errors.Wrap(err, "failed to query_range prometheus")
	}

	defer p.Close()

	return response, nil
}

// IsMetricsExist checks if metrics are existing in Prometheus,
// if there is an error while quering to Prometheus, returns (false, nil, error)
// if there are metrics not existing in Prometheus, returns (false, nonExistMetrics, nil)
// or if all metrics are existing, returns (true, nil, nil)
func (p *Prometheus) IsMetricsExist(ctx context.Context, metrics []string) (bool, []string, error) {
	defer p.Close()

	type result struct {
		metric string
		exist  bool
	}
	ch := make(chan result)
	wg := errgroup.Group{}
	for _, metric := range metrics {
		copyMetric := metric
		wg.Go(func() error {
			resp, err := p.Query(ctx, copyMetric, nil, nil)
			if err != nil {
				return err
			}
			entities, err := resp.GetEntities()
			if err != nil {
				return err
			}
			exist := false
			if len(entities) > 0 {
				exist = true
			}
			result := result{
				metric: copyMetric,
				exist:  exist,
			}
			ch <- result
			return nil
		})
	}

	metricsNotExist := []string{}
	done := make(chan bool)
	go func() {
		for result := range ch {
			if !result.exist {
				metricsNotExist = append(metricsNotExist, result.metric)
			}
		}
		done <- true
	}()

	err := wg.Wait()
	close(ch)
	<-done
	if err != nil {
		return false, nil, errors.Wrap(err, "")
	}

	if len(metricsNotExist) > 0 {
		return false, metricsNotExist, nil
	}
	return true, nil, nil
}

// Close free resource used by prometheus
func (p *Prometheus) Close() {
	p.Transport.CloseIdleConnections()
}