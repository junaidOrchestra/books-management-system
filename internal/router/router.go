package router

import (
	"books-management-system/internal/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	Controllers []controllers.Controller
}

// NewRouter initializes the router with controllers
func NewRouter(controller []controllers.Controller) *Router {
	r := &Router{
		Engine:      gin.Default(),
		Controllers: controller,
	}
	r.setupRoutes()
	return r
}

func (r *Router) setupRoutes() {
	for _, controller := range r.Controllers {
		controller.InitRoutes(r.Engine)
	}
}

// StartServer runs the Gin server
func (r *Router) StartServer() {
	r.Engine.Run(":8080")
}
