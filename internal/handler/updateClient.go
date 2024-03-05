package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateClientHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Update Client",
	})
}
