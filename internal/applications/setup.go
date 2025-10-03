package applications

import (
	"context"
	"os"

	"github.com/cuonglv-smartosc/golang-boiler-template/pkg/cache"
	"github.com/go-redis/redis/v8"

	log "github.com/sirupsen/logrus"

	"github.com/cuonglv-smartosc/golang-boiler-template/docs"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/config"
	"github.com/cuonglv-smartosc/golang-boiler-template/internal/repository/postgres"
	"github.com/cuonglv-smartosc/golang-boiler-template/pkg/sentry"
	"github.com/cuonglv-smartosc/golang-boiler-template/pkg/viper"
)

const defaultConfigPath = "config.yml"

func InitConfig() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = defaultConfigPath
	}

	viper.Load(path, &config.Default)
}

func InitLogging() {
	logLevel, err := log.ParseLevel(config.Default.LogLevel)
	if err != nil {
		log.WithError(err).Fatal("Log level parsing error")
	}

	log.SetLevel(logLevel)
}

func InitSentry() {
	if err := sentry.SetupSentry(
		config.Default.Sentry.DSN,
		sentry.WithSampleRate(config.Default.Sentry.SampleRate),
	); err != nil {
		log.WithError(err).Fatal("Sentry init error")
	}
}

func InitDatabase() {
	db, err := postgres.New(config.Default.Database.URL, config.Default.Database.Log)
	if err != nil {
		log.WithError(err).Fatal("Database init error")
	}

	if err := postgres.Setup(db); err != nil {
		log.WithError(err).Fatal("Database setup error")
	}
}

func InitRedis() {
	cache.RedisClient = &cache.Cache{redis.NewClient(&redis.Options{
		Addr:     config.Default.Redis.URL,
		Password: config.Default.Redis.Password,
	})}
	_, err := cache.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("has error occur when connect to redis: %v", err)
	}
}

func InitSwaggerInfo() {
	docs.SwaggerInfo.Host = config.Default.Swagger.Hostname
}
