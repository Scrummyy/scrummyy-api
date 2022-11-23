package middleware

import (
	"github.com/Scrummyy/scrummyy-api/api/controllers"
	"github.com/gin-gonic/gin"
)

var auth = new(controllers.AuthController)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.TokenValid(c)
		c.Next()
	}
}

// import (
// 	"github.com/Scrummyy/scrummyy-api/internal/constants"
// 	"acquia/decision-service/internal/logger_messages"
// 	"acquia/decision-service/internal/output"
// 	apperror "acquia/decision-service/pkg/app_error"
// 	"acquia/decision-service/pkg/authorization"
// 	"acquia/decision-service/pkg/cache"
// 	"acquia/decision-service/pkg/logger"
// 	"acquia/decision-service/pkg/rest"
// 	"acquia/decision-service/tools"
// 	"errors"
// 	"github.com/gin-gonic/gin"
// 	"github.com/spf13/viper"
// 	"net/http"
// )

// // Authorization handler authorize incoming requests.
// // ******** TODO: Revisit error responses to hide internal errors.
// func Authorization(conf *viper.Viper) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var (
// 			hc, ca interface{}
// 			ok     bool
// 		)

// 		// If the path is part of the exclusion list, skip authorization.
// 		match, err := tools.IsExcludedPath(conf.GetStringMapStringSlice(constants.AuthorizationExcludePaths), c.Request.Method, c.Request.URL.Path)
// 		if err != nil {
// 			logger.GlobalLogger(c).Errorf(logger_messages.ErrorAuthentication, err)
// 			output.Error(c, apperror.New(err, http.StatusInternalServerError))
// 			c.Abort()
// 			return
// 		}
// 		if match {
// 			logger.GlobalLogger(c).Info(logger_messages.InfoAuthorizationSkipped)
// 			return
// 		}

// 		if hc, ok = c.Get(constants.KeyHttpClientInstance); !ok {
// 			logger.GlobalLogger(c).Errorf(logger_messages.ErrorAuthentication, constants.ErrorMissingHTTPClient)
// 			output.Error(c, apperror.New(errors.New(constants.ErrorMissingHTTPClient), http.StatusInternalServerError))
// 			c.Abort()
// 			return
// 		}

// 		if ca, ok = c.Get(constants.KeyCacheInstance); !ok {
// 			logger.GlobalLogger(c).Errorf(logger_messages.ErrorAuthentication, constants.ErrorMissingCacheInstance)
// 			output.Error(c, apperror.New(errors.New(constants.ErrorMissingCacheInstance), http.StatusInternalServerError))
// 			c.Abort()
// 			return
// 		}

// 		// account id is mandatory for authorization.
// 		accountID := c.Param(constants.ParamAccountID)
// 		if accountID == "" {
// 			logger.GlobalLogger(c).Error(constants.ErrorMissingAccountID)
// 			output.Error(c, apperror.New(errors.New(constants.ErrorMissingAccountID), http.StatusUnauthorized))
// 			c.Abort()
// 			return
// 		}

// 		auth, err := authorization.NewAuthorizer(c, hc.(rest.HttpClientInterface), ca.(cache.Storer), conf)
// 		if err != nil {
// 			logger.GlobalLogger(c).Error(err)
// 			output.Error(c, apperror.New(err, http.StatusInternalServerError))
// 			c.Abort()
// 			return
// 		}

// 		if err = auth.PreHook(c); err != nil {
// 			logger.GlobalLogger(c).Error(err)
// 			output.Error(c, apperror.New(err, http.StatusUnauthorized))
// 			c.Abort()
// 			return
// 		}

// 		if err = auth.Authorize(c); err != nil {
// 			logger.GlobalLogger(c).Error(err)
// 			output.Error(c, apperror.New(err, http.StatusUnauthorized))
// 			c.Abort()
// 			return
// 		}
// 	}
// }
