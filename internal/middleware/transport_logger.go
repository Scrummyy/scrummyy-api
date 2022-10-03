package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/Scrummyy/scrummyy-api/internal/constants"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// TransportLogger add middleware to log incoming requests.
func TransportLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		err := c.Request.Body.Close() //  must close
		if err != nil {
			logrus.Error(err)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		logrus.WithField(constants.HeaderRequestID, c.GetString(constants.HeaderRequestID)).
			WithField("request_body", string(bodyBytes)).
			WithField("path", c.Request.RequestURI).
			WithField("http_method", c.Request.Method).Info("http request")
	}
}
