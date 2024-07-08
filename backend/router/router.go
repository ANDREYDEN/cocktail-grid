package router

import (
	scope "cocktail-grid/backend/auth"
	controllers "cocktail-grid/backend/controllers"
	"cocktail-grid/backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"https://cocktailgrid.com",
		"http://cocktailgrid.com",
		"https://127.0.0.1:3000",
		"http://127.0.0.1:3000",
		"http://localhost:3000",
		"http://159.89.80.154",
		"https://159.89.80.154",
	}
	corsConfig.AllowHeaders = []string{
		"Authorization",
		"Content-Type",
	}
	router.Use(cors.New(corsConfig))

	cocktailsGroup := router.Group("/cocktails")
	{
		cocktailController := new(controllers.CocktailController)

		cocktailsGroup.GET("", cocktailController.GetAllCocktails)
		cocktailsGroup.POST(
			"",
			middleware.EnsureValidToken(scope.CreateCocktail),
			cocktailController.CreateCocktail,
		)
		cocktailsGroup.PUT(
			"",
			middleware.EnsureValidToken(scope.UpdateCocktail),
			cocktailController.UpdateCocktail,
		)

		cocktailGroup := cocktailsGroup.Group("/:cocktailId")
		{
			cocktailGroup.DELETE(
				"",
				middleware.EnsureValidToken(scope.DeleteCocktail),
				cocktailController.DeleteCocktail,
			)

			ingredientsGroup := cocktailGroup.Group("/ingredients")
			{
				ingredientGroup := ingredientsGroup.Group("/:ingredientId")
				{
					cocktailIngredientController := new(controllers.CocktailIngredientController)

					ingredientGroup.POST(
						"",
						middleware.EnsureValidToken(scope.CreateCocktailIngredient),
						cocktailIngredientController.CreateCocktailIngredient,
					)
					ingredientGroup.DELETE(
						"",
						middleware.EnsureValidToken(scope.DeleteCocktailIngredient),
						cocktailIngredientController.DeleteCocktailIngredient,
					)
					ingredientGroup.PUT(
						"",
						middleware.EnsureValidToken(scope.UpdateCocktailIngredient),
						cocktailIngredientController.UpdateCocktailIngredient,
					)
				}
			}
		}
	}

	ingredientsGroup := router.Group("/ingredients")
	{
		ingredientController := new(controllers.IngredientController)

		ingredientsGroup.GET("", ingredientController.GetAllIngredients)
		ingredientsGroup.POST(
			"",
			middleware.EnsureValidToken(scope.CreateIngredient),
			ingredientController.CreateIngredient,
		)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
