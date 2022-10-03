package logger

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type LoggerTransport struct {
	http.Transport
}

func (lt *LoggerTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	logRequest(request)
	response, err := lt.Transport.RoundTrip(request)
	logResponse(response)

	return response, err
}

// logRequest logs the details of an outgoing request.
//  ******* TODO: Get request id to connect outgoing requests to their sources.
func logRequest(request *http.Request) {
	if request != nil {
		entry := logrus.WithField("host", request.Host).
			WithField("path", request.URL.Path).
			WithField("http_method", request.Method)

		if request.Body != nil {
			bodyBytes, _ := ioutil.ReadAll(request.Body)
			err := request.Body.Close() //  must close
			if err != nil {
				logrus.Error(err)
			}
			request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			entry.WithField("request_body", string(bodyBytes))
		}
		entry.Debug("outgoing http request")
	}
	logrus.Debug("empty request")
}

// logResponse logs the details of the response.
//  ******* TODO: Get request id to connect outgoing requests to their sources.
func logResponse(response *http.Response) {
	if response != nil {
		entry := logrus.WithField("status_code", response.StatusCode)

		if response.Body != nil {
			bodyBytes, _ := ioutil.ReadAll(response.Body)
			err := response.Body.Close() //  must close
			if err != nil {
				logrus.Error(err)
			}
			response.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			entry.WithField("response_body", string(bodyBytes)).Info("incoming http response")
		}
		entry.Debug("http response")
		return
	}
	logrus.Debug("empty response")
}
