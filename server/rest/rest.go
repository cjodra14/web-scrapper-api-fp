package rest

import (
	"fmt"
	"web-scrapper-api-fp/configuration"
	"web-scrapper-api-fp/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/open-feature/go-sdk/openfeature"
)

func Init(config configuration.RestConfiguration, scrapperService services.ScrapperService, openFeatureClient *openfeature.Client) error {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/healthz"},
	}))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	router.Use(cors.New(corsConfig))

	scrapperGroup := router.Group("/")
	ScrapperRouter(scrapperGroup, scrapperService, openFeatureClient)

	serverAddress := fmt.Sprintf("%s:%s", config.Host, config.Port)

	return router.Run(serverAddress)
}
