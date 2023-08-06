package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"st_api/pkg/handler"
	"st_api/pkg/repository"
	"st_api/pkg/server"
	"st_api/pkg/service"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title           SunTechnics API - OpenAPI
// @version         1.0
// @description     REST API for ecology section of arctech DATA

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	f := setLogger()
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("PG_HOSTNAME"),
		Port:     os.Getenv("PG_PORT"),
		Username: os.Getenv("PG_USER"),
		DBName:   os.Getenv("PG_DBNAME"),
		SSLMode:  os.Getenv("PG_SSLMODE"),
		Password: os.Getenv("PG_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes(viper.GetString("port"))); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("SunTechnics API App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("SunTechnics API Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
	f.Close()
}

func setLogger() *os.File {
	logFile := viper.GetString("log.File")
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFile)
		panic(err)
	}
	logrus.SetOutput(f)
	logrus.SetLevel(logrus.InfoLevel)
	gin.DefaultWriter = io.MultiWriter(f)
	return f
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
