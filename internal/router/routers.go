package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine) {

	v1 := r.Group("api/v1")

	{
		v1.GET("/client", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Get Client",
			})
		})
		v1.POST("/client")
		v1.DELETE("/client")
		v1.PUT("/client")
		v1.GET("/clients")

	}
}
