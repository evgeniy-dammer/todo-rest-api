package main

import (
	todo "github.com/evgeniy-dammer/todo-rest-api"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/handler"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/repository"
	"github.com/evgeniy-dammer/todo-rest-api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

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
		log.Fatalf("failed to initializer database: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
