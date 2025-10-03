package routers

import (
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/middleware"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/modules/auth/controllers"
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
	middleware := middleware.NewAuthMiddleware()

	router.GET("/health", r.controller.HealthCheck)
	router.POST("/login", r.controller.Login)
	router.POST("/register", r.controller.Register)
	router.POST("/refresh", r.controller.RefreshToken)
	router.GET("/me", middleware.Authenticate(), r.controller.GetMe)
}
