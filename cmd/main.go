package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"regulations-api/models"
	"regulations-api/pkg/handler"
	"regulations-api/pkg/repository"
	"regulations-api/pkg/service"
	"regulations-api/pkg/usecase"
	"regulations-api/server"
	"strings"
	"syscall"
	"time"
)

const (
	envFilePath    = `.env`
	configFilePath = "configs/config.json"
)

// @BasePath  		/api/
func main() {
	logrus.Info("start server run")

	environment := &models.Environment{}
	configService := &models.ConfigService{}

	if err := runLogger(); err != nil {
		logrus.Fatal(err.Error())
	}

	if err := loadEnvironment(environment); err != nil {
		logrus.Fatalf(err.Error())
		return
	}
	if err := loadConfig(configService); err != nil {
		logrus.Fatalf(err.Error())
		return
	}
	logrus.Info("load local configs success")

	var serverInstance server.Server

	businessDatabase := repository.NewRegulationsDatabase(*configService, *environment)

	repositorySources := repository.Sources{
		BusinessDB: businessDatabase,
	}
	repos := repository.NewRepository(&repositorySources)
	services := service.NewService(repos, configService)

	usecases := usecase.NewUsecase(services)
	handlers := handler.NewHandler(usecases, services)

	go runServer(&serverInstance, handlers, &configService.Server)

	runChannelStopServer()

	serverInstance.Shutdown(businessDatabase, context.Background())
}

func runServer(server *server.Server, handlers *handler.Handler, config *models.ServerConfig) {
	ginEngine := handlers.InitHTTPRoutes(config)

	if err := server.Run(config.Port, ginEngine); err != nil {
		if err.Error() != "http: Server closed" {
			logrus.Fatalf("error occurred while running http server: %s", nil, err.Error())
		}
	}
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	prefixPath := pwd + "/"

	shortFilePath := strings.TrimPrefix(filepath.ToSlash(entry.Caller.File), filepath.ToSlash(prefixPath))

	var fields string
	for key, value := range entry.Data {
		fields += fmt.Sprintf("\"%s\":\"%v\",", key, value)
	}

	if len(fields) > 0 {
		fields = fields[:len(fields)-1]
	}

	if len(fields) > 0 {
		fields = ", " + fields
	}

	log := fmt.Sprintf(
		"{\"level\":\"%s\",\"msg\":\"%s\",\"point\": \" %s:%d \",\"short_point\":\"%s:%d\", \"time\":\"%s\"%s}\n",
		entry.Level.String(),
		entry.Message,
		entry.Caller.File,
		entry.Caller.Line,
		shortFilePath,
		entry.Caller.Line,
		entry.Time.Format(time.RFC3339),
		fields,
	)
	return []byte(log), nil
}

func runLogger() error {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&CustomFormatter{})

	currentTime := time.Now()
	yearMonthDir := fmt.Sprintf("logs/%d-%02d", currentTime.Year(), currentTime.Month())

	err := os.MkdirAll(yearMonthDir, os.ModePerm)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	logFile := fmt.Sprintf("%s/%d-%02d-%02d.log", yearMonthDir, currentTime.Year(), currentTime.Month(), currentTime.Day())

	logFileHandle, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	logrus.SetOutput(io.MultiWriter(os.Stdout, logFileHandle))

	return nil
}

func loadEnvironment(environment *models.Environment) error {
	if err := godotenv.Load(envFilePath); err != nil {
		logrus.Warning("load file not found, Environment variables load from Environment")
	}
	if err := env.Parse(environment); err != nil {
		return err
	}

	return nil
}

func loadConfig(config *models.ConfigService) error {
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}

	return nil
}

func runChannelStopServer() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
}
