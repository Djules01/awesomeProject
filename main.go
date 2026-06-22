package main

import (
	httpadapter "awesomeProject/adapters/http"
	mongodbadapter "awesomeProject/adapters/mongodb"
	translationadapter "awesomeProject/adapters/translation"
	"awesomeProject/configuration"
	"awesomeProject/service"
	"context"
	"fmt"
	"net/http"
)

func main() {
	config := configuration.LoadConfig()

	err := run(config, http.ListenAndServe)
	if err != nil {
		panic(err)
	}
}

func run(config configuration.Config, listenAndServe func(string, http.Handler) error) error {
	client, err := mongodbadapter.InitClient(config.MongoURI)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	translator := translationadapter.NewTranslationService(10)
	translator.StartWorker()

	todoRepository := mongodbadapter.NewRepository(client, config.MongoDB)
	todoService := service.NewService(todoRepository)
	todoHandler := httpadapter.NewHandler(todoService)
	r := httpadapter.NewRouter(todoHandler, config.APIKey)

	fmt.Println("Serveur lance sur http://localhost:8080")

	err = listenAndServe(":"+config.Port, r)
	if err != nil {
		return err
	}

	return nil
}
