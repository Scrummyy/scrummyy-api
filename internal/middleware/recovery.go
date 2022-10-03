package middleware

import (
	"acquia/decision-service/internal/logger_messages"
	"acquia/decision-service/internal/output"
	apperror "acquia/decision-service/pkg/app_error"
	"errors"
	"net/http"

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
