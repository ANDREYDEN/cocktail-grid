package main

import (
	scope "cocktail-grid/backend/auth"
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/docs"
	"cocktail-grid/backend/router"
	"cocktail-grid/backend/seeders"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//	@title						CocktailGrid API
//	@description				An API to get cocktails and ingredients
//	@version					1
//	@BasePath					/
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	halt := handleArguments()
	if halt {
		return
	}

	docs.SwaggerInfo.BasePath = "/"

	db.Init()
	scope.InitScopes()
	router := router.NewRouter()
	port := "8080"
	log.Printf("\nVisit http://localhost:%v/swagger/index.html to view Swagger documentation.\n\n", port)
	router.Run(fmt.Sprintf("0.0.0.0:%v", port))
}

func handleArguments() bool {
	args := os.Args[1:]

	if len(args) == 0 {
		return false
	}

	switch args[0] {
	case "--seed":
		seeders.Seed()
		return true
	}

	return false
}
