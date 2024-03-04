package router

import (
	"github.com/LukasLimalkl/trafego-back/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine) {

	v1 := r.Group("api/v1")

	{
		v1.GET("/client", handler.GetClientHandler)

		v1.POST("/client", handler.CreateClientHandler)
		v1.DELETE("/client", handler.DeleteClientHandler)
		v1.PUT("/client", handler.UpdateClientHandler)
		v1.GET("/clients", handler.ListClientHandler)

	}
}
