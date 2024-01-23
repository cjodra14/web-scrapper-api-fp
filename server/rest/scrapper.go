package rest

import (
	"web-scrapper-api-fp/handlers"
	"web-scrapper-api-fp/handlers/middlewares"
	"web-scrapper-api-fp/services"

	"github.com/gin-gonic/gin"
	"github.com/open-feature/go-sdk/openfeature"
)

func ScrapperRouter(router *gin.RouterGroup, scrapperService services.ScrapperService, openFeatureClient *openfeature.Client) {
	router.GET("/scrap", middlewares.FeatureToggleMiddleware(openFeatureClient, "scrap"), handlers.GetURIScrapped(scrapperService))
	router.GET("/healthz", handlers.GETHealthz())
}
