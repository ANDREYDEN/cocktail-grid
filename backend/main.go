package main

import (
	docs "cocktail-grid/backend/docs"
)

// @title CocktailGrid API
// @description An API to get cocktails and ingredients
// @version 1
// @BasePath /

func main() {
	docs.SwaggerInfo.BasePath = "/"

	router := NewRouter()
	router.Run("0.0.0.0:8080")
}
