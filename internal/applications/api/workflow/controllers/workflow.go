package controllers

import (
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/workflow/handlers"
	"net/http"

	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"
	"github.com/gin-gonic/gin"
)

type WorkflowController struct {
	service *handlers.WorkflowService
}

func NewWorkflowController(db repository.Storage) *WorkflowController {
	return &WorkflowController{service: handlers.NewWorkflowService(db)}
}

func (controller *WorkflowController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
