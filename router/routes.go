package router

import (
	"github.com/gin-gonic/gin"
	"github.com/imatheus-lucas/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunity", handler.ShowOpeningHandler)
		v1.POST("/opportunities", handler.CreateOpeningHandler)
		v1.DELETE("/opportunities", handler.DeleteOpeningHandler)
		v1.PUT("/opportunities", handler.UpdateOpeningHandler)
		v1.GET("/opportunities", handler.ListOpeningHandler)
	}
}
