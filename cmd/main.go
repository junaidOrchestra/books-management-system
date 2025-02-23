package main

import (
	_ "books-management-system/docs"
	_ "books-management-system/internal/controllers"
	"books-management-system/internal/router"
	"books-management-system/modules"
	"books-management-system/utils"
	"go.uber.org/fx"
)

func main() {
	utils.InitLogger()
	app := fx.New(
		modules.Module,
		fx.Invoke(registerRoutes), // Automatically registers routes
	)

	app.Run() // Start the application
}

// Registers API routes using Fx
func registerRoutes(r *router.Router) {
	go r.StartServer() // Start Gin server
}
