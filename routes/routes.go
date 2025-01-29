package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matdorneles/gestao_ti/handlers"
)

// run web server and handle http requests
func HandleRequests() {
	r := gin.Default()

	// COMPUTER ROUTES
	computer := r.Group("/computer")
	{
		computer.GET("/", handlers.GetAllComputers)
		computer.GET("/:id", handlers.GetComputer)
		computer.POST("/", handlers.CreateNewComputer)
		computer.PATCH("/", handlers.PatchComputer)
		computer.DELETE("/:id", handlers.DeleteComputer)
	}

	r.Run() // default port :8080
}
