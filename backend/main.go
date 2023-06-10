package main

import (
	docs "cocktail-grid/backend/docs"
	"fmt"
)

// @title CocktailGrid API
// @description An API to get cocktails and ingredients
// @version 1
// @BasePath /

func main() {
	docs.SwaggerInfo.BasePath = "/"

	port := "8080"
	router := NewRouter()
	fmt.Printf("\nVisit http://localhost:%v/swagger/index.html to view Swagger documentation.\n\n", port)
	router.Run(fmt.Sprintf("0.0.0.0:%v", port))
}
