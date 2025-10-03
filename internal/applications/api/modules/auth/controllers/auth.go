package controllers

import (
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/errors"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/modules/auth/dtos"

	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/modules/auth/handlers"
	log "github.com/sirupsen/logrus"

	"strconv"

	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"
	resp "github.com/cuonglv-smartosc/golang-boiler-template/pkg/response"
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	HealthCheck(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	GetMe(c *gin.Context)
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

func (authController *AuthController) Login(c *gin.Context) {
	var req dtos.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid request body: ", err)
		if errors.HandleValidationError(c, err) {
			return
		}
	}

	response, err := authController.service.Login(&req)
	if err != nil {
		log.Error("Login failed: ", err)
		errors.HandleCustomError(c, err, "", "")
		return
	}

	resp.OK(c, response)
}

func (authController *AuthController) RefreshToken(c *gin.Context) {
	var req dtos.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Invalid request body: ", err)
		if errors.HandleValidationError(c, err) {
			return
		}
	}

	response, err := authController.service.RefreshToken(&req)
	if err != nil {
		log.Error("Token refresh failed: ", err)
		errors.HandleCustomError(c, err, "", "")
		return
	}

	resp.OK(c, response)
}

func (authController *AuthController) GetMe(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		errors.HandleCustomError(c, nil, "UNAUTHORIZED", "Missing user id in token")
		return
	}

	userIDStr, ok := userIDVal.(string)
	if !ok {
		errors.HandleCustomError(c, nil, "UNAUTHORIZED", "Invalid user id type in token")
		return
	}

	userID, convErr := strconv.ParseInt(userIDStr, 10, 64)
	if convErr != nil {
		errors.HandleCustomError(c, nil, "UNAUTHORIZED", "Invalid user id in token")
		return
	}

	info, err := authController.service.GetUserByID(userID)
	if err != nil {
		errors.HandleCustomError(c, err, "", "")
		return
	}

	resp.OK(c, info)
}
