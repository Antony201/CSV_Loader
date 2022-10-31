package main

import (
	loader "github.com/Antony201/CsvLoader"
	"github.com/Antony201/CsvLoader/pkg/handler"
	"github.com/Antony201/CsvLoader/pkg/repository"
	"github.com/Antony201/CsvLoader/pkg/service"
	"io"
	"os"
	"os/signal"
	"syscall"

	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title CSVLoader App API
// @version 1.0
// @description API server for CSV Application

// @host docker_host:8000
// @BasePath /

func initLogger() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logFile := "/logs/api.log"

	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logrus.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	logrus.SetOutput(mw)
}
func main() {
	initLogger()

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing envs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("initializing db is failed: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(loader.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("CsvLoader Started!")

	quit := make(chan os.Signal, 1) // lock main function
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Print("CsvLoader Shutting Down!")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}