package routers

import (
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/auth/controllers"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	controller *controllers.AuthController
}

func NewAuthRouter(db repository.Storage) *AuthRouter {
	return &AuthRouter{
		controller: controllers.NewAuthController(db),
	}
}

func (r *AuthRouter) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/health", r.controller.HealthCheck)
}
