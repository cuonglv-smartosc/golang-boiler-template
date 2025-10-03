package routers

import (
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/workflow/controllers"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, db repository.Storage) {
	apiRouteGroup := r.Group("/workflow")

	workflowController := controllers.NewWorkflowController(db)
	apiRouteGroup.GET("/health", workflowController.HealthCheck)
}
