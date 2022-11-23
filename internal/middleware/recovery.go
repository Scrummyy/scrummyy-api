package middleware

import (
	"errors"
	"net/http"

	apperror "github.com/Scrummyy/scrummyy-api/internal/app_error"
	"github.com/Scrummyy/scrummyy-api/internal/logger_messages"
	"github.com/Scrummyy/scrummyy-api/internal/output"

	"github.com/gin-gonic/gin"
)

// CustomRecovery adds a custom middleware to recover from panic.
func CustomRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if _, ok := recovered.(string); ok {
			output.Error(c, apperror.New(recovered.(error), http.StatusInternalServerError))
		}
		output.Error(c, apperror.New(errors.New(logger_messages.InternalServerError), http.StatusInternalServerError))
	})
}
