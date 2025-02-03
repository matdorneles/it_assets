package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matdorneles/gestao_ti/api/pkg/computers"
	"github.com/matdorneles/gestao_ti/api/pkg/models"
)

// receives a call from router and registers the routes of Computer model
func computerRoutes(r *gin.Engine) {
	// Instantiate computer Handler and provide a data computer implementation
	computer := models.NewComputer()
	ch := computers.NewComputerHandler(computer)

	computerGroup := r.Group("/computer")
	{
		computerGroup.GET("/", ch.List)
		computerGroup.GET("/:id", ch.Get)
		computerGroup.POST("/", ch.Create)
		computerGroup.PATCH("/:id", ch.Update)
		computerGroup.DELETE("/:id", ch.Delete)
	}
}
