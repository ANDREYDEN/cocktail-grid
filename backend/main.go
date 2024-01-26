package main

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/docs"
	"cocktail-grid/backend/seeders"
	"fmt"
	"log"
	"os"
)

//	@title			CocktailGrid API
//	@description	An API to get cocktails and ingredients
//	@version		1
//	@BasePath		/
func main() {
	halt := handleArguments()
	if (halt) {
		return
	}

	docs.SwaggerInfo.BasePath = "/"

	db.Init()

	port := "8080"
	router := NewRouter()
	log.Printf("\nVisit http://localhost:%v/swagger/index.html to view Swagger documentation.\n\n", port)
	router.Run(fmt.Sprintf("0.0.0.0:%v", port))
}

func handleArguments() bool {
	args := os.Args[1:]

	if (len(args) == 0) {
		return false
	}
	
	switch args[0] {
		case "--seed":
			seeders.Seed()
			return true
	}

	return false
}