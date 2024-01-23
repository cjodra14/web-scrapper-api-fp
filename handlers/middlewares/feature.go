package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/open-feature/go-sdk/openfeature"
)

const defaultFeatureFlag = false

func FeatureToggleMiddleware(client *openfeature.Client, featureName string) gin.HandlerFunc {
	return func(c *gin.Context) {

		isFeatureEnabled, err := client.BooleanValue(c, featureName, defaultFeatureFlag, openfeature.EvaluationContext{})
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "error evaluating feature availability"})
			return
		}

		if !isFeatureEnabled {
			c.AbortWithStatusJSON(403, gin.H{"error": "feature not enabled"})
			return
		}

		c.Next()
	}
}
