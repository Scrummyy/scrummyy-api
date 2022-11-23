package server

import (
	"io"

	db "github.com/Scrummyy/scrummyy-api/configs"
	"github.com/Scrummyy/scrummyy-api/internal/constants"

	"github.com/Scrummyy/scrummyy-api/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sebest/xff"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetRouter initializes the framework router.
func GetRouter(conf *viper.Viper) *gin.Engine {
	router := gin.New()

	// Disable debug mode of Gin.
	gin.SetMode(gin.ReleaseMode)

	// Discard DefaultWriter of Gin to prevent it from logging.
	gin.DefaultWriter = io.Discard

	// handle unknown routes
	router.NoRoute(func(c *gin.Context) {
		logrus.Debugf("%s request to unknown path %s (referer: %s)", c.Request.Method, c.Request.URL.Path, c.Request.Header.Get("referer"))
	})

	//**********register middlewares**********

	// Recovery from panic.
	router.Use(middleware.CustomRecovery())

	// Add request id to incoming requests.
	router.Use(middleware.AssignRequestID())

	// Add logger for http request
	router.Use(middleware.TransportLogger())

	// Add CORS middleware
	if conf.GetBool(constants.IsCorsEnabled) {
		router.Use(cors.New(cors.Config{
			AllowAllOrigins:  conf.GetBool(constants.CorsAllowAllOrigins),
			AllowHeaders:     conf.GetStringSlice(constants.CorsAllowedHeaders),
			AllowMethods:     conf.GetStringSlice(constants.CorsAllowedMethods),
			AllowCredentials: true,
		}))
	}

	RegisterDependencyMiddleware(router, conf)

	// Add in middleware to get the real client IP-address from the request
	// headers. It must be added before auth middleware.
	x, _ := xff.Default()
	router.Use(func(c *gin.Context) {
		x.HandlerFunc(c.Writer, c.Request)
	})

	// Add authentication middleware.
	// if conf.GetBool(constants.AuthenticationEnabled) {
	// 	router.Use(middleware.Authentication(conf))
	// }

	// // Add authorization middleware.
	// router.Use(middleware.Authorization(conf))

	// Add further middlewares like monitoring, etc

	return router
}

func RegisterDependencyMiddleware(router *gin.Engine, conf *viper.Viper) {
	// add dependecies to request contex
	// ******* TODO: need to come up with a more robust solution to avoid
	// cluttering in the current file
	// I know it's not a cleaner approach but will work on it and refine the approach
	var dependencyMap map[string]interface{}

	dependencyMap = make(map[string]interface{})

	db.Init(conf)

	// hc := InitHttpClient(conf)

	// ch, err := InitCaching(conf)
	// if err != nil {
	// 	logrus.WithError(err).Warning("failed to connect to redis server")
	// }

	// dependencyMap[constants.KeyDecisionDBInstance] = decisionDB
	// dependencyMap[constants.KeyHttpClientInstance] = hc
	// dependencyMap[constants.KeyCacheInstance] = ch

	router.Use(middleware.InjectDependency(&dependencyMap))
}
