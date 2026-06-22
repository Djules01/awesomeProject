package main

import (
	httpadapter "awesomeProject/adapters/http"
	sqliteadapter "awesomeProject/adapters/sqlite"
	translationadapter "awesomeProject/adapters/translation"
	"awesomeProject/configuration"
	"awesomeProject/service"
	"fmt"
	"net/http"
)

func main() {
	config := configuration.LoadConfig()

	db, err := sqliteadapter.InitDB(config.DBPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	translator := translationadapter.NewTranslationService(10)
	translator.StartWorker()

	todoRepository := sqliteadapter.NewRepository(db)
	todoService := service.NewService(todoRepository)
	todoHandler := httpadapter.NewHandler(todoService)
	r := httpadapter.NewRouter(todoHandler, config.APIKey)

	fmt.Println("Serveur lance sur http://localhost:8080")

	err = http.ListenAndServe(":"+config.Port, r)
	if err != nil {
		panic(err)
	}
}
