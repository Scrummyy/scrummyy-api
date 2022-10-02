package server

import (
	"acquia/decision-service/api/v3/campaigns"
	"acquia/decision-service/api/v3/decision"
	"acquia/decision-service/api/v3/pages"
	"acquia/decision-service/api/v3/rules"
	"acquia/decision-service/api/v3/sites"
	"acquia/decision-service/api/v3/slots"
	"acquia/decision-service/api/v3/userinfo"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RegisterRoutes(r *gin.Engine, conf *viper.Viper) {
	// Add Ping EP for health checks
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong!"})
	})

	// v3 routes
	apiV3 := r.Group("/v3")
	apiWithAccountV3 := r.Group("/v3/accounts/:account_id")

	{
		campaigns.RegisterCampaignHandlers(apiV3, conf)
		userinfo.RegisterUserInfoHandlers(apiV3, conf)
	}
	{
		campaigns.RegisterAccountCampaignHandlers(apiWithAccountV3, conf)
		slots.RegisterSlotHandlers(apiWithAccountV3, conf)
		pages.RegisterPageHandlers(apiWithAccountV3, conf)
		rules.RegisterRuleHandlers(apiWithAccountV3, conf)
		sites.RegisterSiteHandlers(apiWithAccountV3, conf)
		decision.RegisterDecisionHandlers(apiWithAccountV3, conf)
	}
}
