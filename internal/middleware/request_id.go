package middleware

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/Scrummyy/scrummyy-api/internal/constants"
	"github.com/twinj/uuid"

	"github.com/gin-gonic/gin"
)

// AssignRequestID adds unique request id to a http request.
func AssignRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string

		headerRequestID := c.Request.Header.Get(constants.HeaderRequestID)
		if headerRequestID != "" {
			requestID = headerRequestID
		} else {
			requestID = generateRequestID()
		}

		c.Set(constants.HeaderRequestID, requestID)
		c.Writer.Header().Set(constants.HeaderRequestIDResponse, requestID)
	}
}

// generateRequestID generates a unique ID.
func generateRequestID() string {
	// seed to get true randomization
	t := uint64(time.Now().UnixNano())
	// perform addition atomic level to avoid collision
	ri, _ := rand.Int(rand.Reader, big.NewInt(100000000))
	atomic.AddUint64(&t, ri.Uint64())
	m := sha256.New()
	m.Write([]byte(fmt.Sprint(t)))
	// make a copy of the original request context
	return hex.EncodeToString(m.Sum(nil))
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}
