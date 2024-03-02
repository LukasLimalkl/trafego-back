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
		v1.POST("/opening")
		v1.DELETE("/opening")
		v1.PUT("/opening")
		v1.GET("/openings")

	}
}
