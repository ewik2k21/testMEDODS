package main

import (
	"log"
	"os"
	testmedods "testMEDODS"
	"testMEDODS/pkg/handler"
	"testMEDODS/pkg/repository"
	"testMEDODS/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("err init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("err env vars:%s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Connection{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DbName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed init db : %s", err.Error())
	}

	rep := repository.NewRepository(db)
	services := service.NewService(rep)
	handlers := handler.NewHandler(services)

	server := new(testmedods.Server)
	if err := server.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
