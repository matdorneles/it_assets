package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matdorneles/gestao_ti/handlers"
)

// receives a call from router and registers the routes of Computer model
func computerRoutes(r *gin.Engine) {
	computer := r.Group("/computer")
	{
		computer.GET("/", handlers.GetAllComputers)
		computer.GET("/:id", handlers.GetComputer)
		computer.POST("/", handlers.CreateNewComputer)
		computer.PATCH("/", handlers.PatchComputer)
		computer.DELETE("/:id", handlers.DeleteComputer)
	}
}
