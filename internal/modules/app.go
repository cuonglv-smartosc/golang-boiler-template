package modules

import (
	"context"

	"github.com/cuonglv-smartosc/golang-boiler-template/pkg/http"
	"github.com/cuonglv-smartosc/golang-boiler-template/pkg/service"
)

type App struct {
	server http.Server
}

func NewApp() *App {
	InitConfig()
	InitLogging()
	InitSentry()
	InitSwaggerInfo()
	InitRedis()

	engine := InitRoutes()
	server := http.NewHTTPServer(engine, "8080")
	return &App{server: server}
}

func (a *App) Run(ctx context.Context) {
	service.RunWithGracefulShutdown(ctx, a.server.Run)
}
