package main

import (
	controllers "cocktail-grid/backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

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

	cocktailsGroup := router.Group("/cocktails")
	{
		cocktailController := new(controllers.CocktailController)

		cocktailsGroup.GET("", cocktailController.GetAllCocktails)
		cocktailsGroup.POST("", cocktailController.CreateCocktail)

		cocktailGroup := cocktailsGroup.Group("/:cocktailId")
		{
			ingredientsGroup := cocktailGroup.Group("/ingredients")
			{
				ingredientGroup := ingredientsGroup.Group("/:ingredientId")
				{
					cocktailIngredientController := new(controllers.CocktailIngredientController)

					ingredientGroup.POST("", cocktailIngredientController.CreateCocktailIngredient)
				}
			}
		}
	}

	ingredientsGroup := router.Group("/ingredients")
	{
		ingredientController := new(controllers.IngredientController)

		ingredientsGroup.GET("", ingredientController.GetAllIngredients)
		ingredientsGroup.POST("", ingredientController.CreateIngredient)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
