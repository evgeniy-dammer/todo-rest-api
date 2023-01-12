package main

import (
	"context"
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/handler"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/repository"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// set log format as JSON
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// initializing config
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	// loading env variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	// establishing database connection
	db, err := repository.NewPostgresDB(
		repository.DbConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			DbName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
	)

	if err != nil {
		logrus.Fatalf("failed to initialize database: %s", err)
	}

	// dependency injections
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// create new server
	srv := new(todo.Server)

	go func() {
		// run server
		if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Println("Application started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Application shutdown...")

	if err = srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error ocured on server shutdown: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error ocured on database connection close: %s", err.Error())
	}
}

// initConfig initializes configuration
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
