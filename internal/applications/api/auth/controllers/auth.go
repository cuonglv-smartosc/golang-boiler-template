package controllers

import (
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/auth/handlers"
	log "github.com/sirupsen/logrus"

	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"
	resp "github.com/cuonglv-smartosc/golang-boiler-template/pkg/response"
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	HealthCheck(c *gin.Context)
}

type AuthController struct {
	service *handlers.AuthService
}

func NewAuthController(db repository.Storage) *AuthController {
	return &AuthController{service: handlers.NewAuthService(db)}
}

func (authController *AuthController) HealthCheck(c *gin.Context) {
	log.Info("Auth Health Check")
	resp.OK(c, map[string]string{"message": "OK"})
}
