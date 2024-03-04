package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClientHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get Client",
	})
}
