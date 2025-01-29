package main

import (
	"github.com/matdorneles/gestao_ti/config"
	"github.com/matdorneles/gestao_ti/routes"
)

func main() {
	config.ConnectDB()
	routes.HandleRequests()
}
