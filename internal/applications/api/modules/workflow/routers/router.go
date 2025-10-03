package routers

import (
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/modules/workflow/controllers"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"
	"github.com/gin-gonic/gin"
)

type WorkflowRouter struct {
	controller *controllers.WorkflowController
}

func NewAuthRouter(db repository.Storage) *WorkflowRouter {
	return &WorkflowRouter{
		controller: controllers.NewWorkflowController(db),
	}
}

func (r *WorkflowRouter) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/health", r.controller.HealthCheck)
}
