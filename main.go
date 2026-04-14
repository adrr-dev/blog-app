package main

import (
	"log"
	"net/http"

	"github.com/adrr-dev/blog-app/components"
	"github.com/adrr-dev/blog-app/internal/database"
	"github.com/adrr-dev/blog-app/internal/handlers"
	"github.com/adrr-dev/blog-app/internal/middleware"
	"github.com/adrr-dev/blog-app/internal/repository"
	"github.com/adrr-dev/blog-app/internal/routes"
	"github.com/adrr-dev/blog-app/internal/service"
)

func main() {
	dataFile := "data.db"
	myDB, err := database.InitializeDB(dataFile)
	if err != nil {
		log.Fatal(err)
	}
	myRepo := repository.NewRepository(myDB)
	myService := service.NewService(myDB, myRepo)
	myMiddleware := middleware.NewMiddleWare()
	myComponents := components.NewComponents()
	myHandling := handlers.NewHandling(myService, myMiddleware, myComponents)
	routes := routes.NewRouter(myHandling, myMiddleware)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	log.Fatal(server.ListenAndServe())
}
