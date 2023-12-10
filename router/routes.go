package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opportunities", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Server is running",
			})
		})
		v1.POST("/opportunities", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Server is running",
			})
		})
		v1.DELETE("/opportunities", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Server is running",
			})
		})
		v1.PUT("/opportunities", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Server is running",
			})
		})
	}
}
