package handlers

import (
	"net/http"
	"web-scrapper-api-fp/services"

	"github.com/gin-gonic/gin"
)

func GetURIScrapped(scrapperService services.ScrapperService) gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Query("uri")

		posts := scrapperService.ScrapePosts(uri)

		c.JSON(http.StatusOK, posts)
	}
}
