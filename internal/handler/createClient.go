package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClientHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Create Client",
	})
}
