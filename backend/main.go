package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type cocktail struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}

type ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type cocktail_ingredient struct {
	CocktailID   int `json:"cocktail_id"`
	IngredientID int `json:"ingredient_id"`
	Quantity     int `json:"quantity"`
}

var cocktails = []cocktail{
	{ID: 1, Title: "Mojito", ImageURL: "https://www.thecocktaildb.com/images/media/drink/3z6xdi1589574608.jpg"},
	{ID: 2, Title: "Margarita", ImageURL: "https://www.thecocktaildb.com/images/media/drink/5noda61589575158.jpg"},
	{ID: 3, Title: "Long Island Iced Tea", ImageURL: "https://www.thecocktaildb.com/images/media/drink/loezxn1504373874.jpg"},
}

func getCocktails(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, cocktails)
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://cocktailgrid.com",
			"https://127.0.0.1:3000",
			"http://127.0.0.1:3000",
		},
	}))

	router.GET("/cocktails", getCocktails)

	router.Run("0.0.0.0:8080")
}
