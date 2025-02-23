package controllers

import "github.com/gin-gonic/gin"

// Controller interface for automatic route registration
type Controller interface {
	InitRoutes(router *gin.Engine)
}
