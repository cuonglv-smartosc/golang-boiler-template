package api

import (
	"context"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/applications"

	"github.com/cuonglv-smartosc/golang-boiler-template/internal/config"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository/postgres"
	"github.com/cuonglv-smartosc/golang-boiler-template/pkg/http"
	"github.com/cuonglv-smartosc/golang-boiler-template/pkg/service"
	log "github.com/sirupsen/logrus"
)

type App struct {
	server http.Server
}

func NewApp() *App {
	applications.InitConfig()
	applications.InitLogging()
	applications.InitSentry()
	applications.InitSwaggerInfo()
	applications.InitRedis()
	applications.InitDatabase()

	db, err := postgres.New(config.Default.Database.URL, config.Default.Database.Log)
	if err != nil {
		log.WithError(err).Fatal("Database init error")
	}

	router := InitRoutes(db)
	server := http.NewHTTPServer(router, config.Default.Port)
	return &App{server: server}
}

func (a *App) Run(ctx context.Context) {
	service.RunWithGracefulShutdown(ctx, a.server.Run)
}
