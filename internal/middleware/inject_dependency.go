package middleware

import (
	"github.com/gin-gonic/gin"
)

// InjectDependency adds dependencies for a map to the gin context which can be accessed in various places like
// controller, model, services etc.
func InjectDependency(dependencies *map[string]interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		for key, inst := range *dependencies {
			c.Set(key, inst)
		}
	}
}
