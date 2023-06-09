package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	docs "cocktail-grid/backend/docs"
	models "cocktail-grid/backend/models"
)

type ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type cocktail_ingredient struct {
	CocktailID   int `json:"cocktail_id"`
	IngredientID int `json:"ingredient_id"`
	Quantity     int `json:"quantity"`
}

// @title CocktailGrid API
// @description An API to get cocktails and ingredients
// @version 1
// @BasePath /

// getCocktails godoc
// @Summary retrieves cocktails
// @Schemes
// @Description do ping
// @Tags cocktails
// @Accept json
// @Produce json
// @Success 200 {object} models.cocktail
// @Router /cocktails [get]
func getCocktails(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.GetAllCocktails())
}

func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://cocktailgrid.com",
			"https://127.0.0.1:3000",
			"http://127.0.0.1:3000",
			"http://localhost:3000",
			"http://159.89.80.154",
			"https://159.89.80.154",
		},
	}))

	router.GET("/cocktails", getCocktails)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("0.0.0.0:8080")
}
