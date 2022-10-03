package logger

import (
	"github.com/Scrummyy/scrummyy-api/internal/constants"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GlobalLogger returns the Global logrus logger instance by attaching the request id fetched from the gin context.
func GlobalLogger(c *gin.Context) *logrus.Entry {
	// attach the request ID to the log if available
	rid := c.GetString(constants.HeaderRequestID)
	if rid != "" {
		return logrus.WithField("request_id", rid)
	}

	return logrus.NewEntry(logrus.StandardLogger())
}
