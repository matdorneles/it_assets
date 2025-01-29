package routes

import (
	"github.com/gin-gonic/gin"
)

// run web server and handle http requests
func HandleRequests() {
	r := gin.Default()

	// computer routes
	computerRoutes(r)

	r.Run() // default port :8080
}
