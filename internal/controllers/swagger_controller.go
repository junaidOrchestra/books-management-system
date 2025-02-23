package controllers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerController struct {
}

func NewSwaggerController() *SwaggerController {
	return &SwaggerController{}
}
func (s *SwaggerController) InitRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
