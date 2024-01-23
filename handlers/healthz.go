package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GETHealthz() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "online")
	}
}