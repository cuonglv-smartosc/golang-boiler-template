package api

import (
	authrouters "github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/modules/auth/routers"
	workflowrouters "github.com/cuonglv-smartosc/golang-boiler-template/internal/applications/api/modules/workflow/routers"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/config"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"
	httplib "github.com/cuonglv-smartosc/golang-boiler-template/pkg/http"
	"github.com/gin-gonic/gin"
)

func InitRoutes(db repository.Storage) *gin.Engine {
	var router *gin.Engine
	if config.Default.Gin.Mode == gin.DebugMode {
		router = gin.Default()
	} else {
		router = gin.New()
	}

	router.Use(httplib.CORSMiddleware())

	apiRouteGroup := router.Group("/api/v1")
	authrouters.NewAuthRouter(db).RegisterRoutes(apiRouteGroup.Group("auth"))
	workflowrouters.NewAuthRouter(db).RegisterRoutes(apiRouteGroup.Group("workflow"))
	return router
}
