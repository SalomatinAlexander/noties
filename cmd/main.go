package main

import (
	"flag"
	"fmt"
	"log"
	"noties/cmd/server"
	"noties/internal/handlers"
	"noties/internal/services"
	"noties/internal/store"

	"github.com/spf13/viper"
)

var (
	configPath string
)

func init() {
	print("INIT VOID")
	flag.StringVar(&configPath, "config-path", "./apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()
	if err := initConfig(); err != nil {
		log.Fatal("ERROR init config:" + fmt.Sprint(err))

	}

	repo := store.NewRepository(store.New(store.NewConfig()))
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	server := server.New(viper.GetString("8080"), handler)
	if err := server.Run(); err != nil {
		log.Fatal("ERROR:" + fmt.Sprint(err))
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}