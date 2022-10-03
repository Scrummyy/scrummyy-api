package middleware

// import (
// 	"acquia/decision-service/internal/constants"
// 	"acquia/decision-service/internal/logger_messages"
// 	"acquia/decision-service/internal/output"
// 	apperror "acquia/decision-service/pkg/app_error"
// 	"acquia/decision-service/pkg/authentication"
// 	"acquia/decision-service/pkg/cache"
// 	"acquia/decision-service/pkg/logger"
// 	"acquia/decision-service/pkg/rest"
// 	"acquia/decision-service/tools"
// 	"errors"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/spf13/viper"
// )

// // Authentication handler authenticates incoming requests.
// // ******** TODO: Revisit error responses to hide internal errors.
// func Authentication(conf *viper.Viper) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var (
// 			hc, ca interface{}
// 			ok     bool
// 		)

// 		// Create a signable response writer to be able to sign the response at the end.
// 		erw := rest.GetSignableResponseWriter(c.Writer)
// 		defer erw.Close()

// 		c.Writer = erw

// 		// If the path is part of the exclusion list, skip authentication.
// 		match, err := tools.IsExcludedPath(conf.GetStringMapStringSlice(constants.AuthenticationExcludePaths), c.Request.Method, c.Request.URL.Path)
// 		if err != nil {
// 			logger.GlobalLogger(c).Errorf(logger_messages.ErrorAuthentication, err)
// 			output.Error(c, apperror.New(err, http.StatusInternalServerError))
// 			c.Abort()
// 			return
// 		}
// 		if match {
// 			logger.GlobalLogger(c).Info(logger_messages.InfoAuthenticationSkipped)
// 			c.Next()
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

// 		authSrv, err := authentication.NewAuthenticator(c, hc.(rest.HttpClientInterface), ca.(cache.Storer), conf)
// 		if err != nil {
// 			logger.GlobalLogger(c).Errorf(logger_messages.ErrorAuthentication, err)
// 			output.Error(c, apperror.New(err, http.StatusInternalServerError))
// 			c.Abort()
// 			return
// 		}

// 		// Pre-authentication tasks
// 		if err = authSrv.PreHook(c); err != nil {
// 			logger.GlobalLogger(c).Errorf(logger_messages.ErrorAuthentication, err)
// 			output.Error(c, apperror.New(err, http.StatusBadRequest))
// 			c.Abort()
// 			return
// 		}

// 		valid, err := authSrv.Authenticate(c)
// 		if err != nil {
// 			logger.GlobalLogger(c).Errorf(logger_messages.ErrorAuthentication, err)
// 			output.Error(c, apperror.New(err, http.StatusInternalServerError))
// 			c.Abort()
// 			return
// 		}
// 		if !valid {
// 			logger.GlobalLogger(c).Errorf(constants.ErrorAuthenticationFailed)
// 			output.Error(c, apperror.New(errors.New(constants.ErrorAuthenticationFailed), http.StatusForbidden))
// 			c.Abort()
// 			return
// 		}

// 		// Run the next handler.
// 		c.Next()

// 		// Post authentication tasks
// 		if err = authSrv.PostHook(c); err != nil {
// 			logger.GlobalLogger(c).Errorf(logger_messages.ErrorAuthentication, err)
// 			output.Error(c, apperror.New(err, http.StatusInternalServerError))
// 			c.Abort()
// 			return
// 		}
// 	}
// }
