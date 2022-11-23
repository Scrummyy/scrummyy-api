package server

import (
	"github.com/Scrummyy/scrummyy-api/api/controllers"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RegisterRoutes(r *gin.Engine, conf *viper.Viper) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
	})

	apiWithAccountV1 := r.Group("/v1")

	{
		controllers.RegisterUserHandler(apiWithAccountV1, conf)
	}
}
