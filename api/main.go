package main

import (
	"github.com/matdorneles/gestao_ti/api/pkg/models"
	"github.com/matdorneles/gestao_ti/api/routes"
)

func main() {
	models.ConnectDB()
	routes.HandleRequests()
}
