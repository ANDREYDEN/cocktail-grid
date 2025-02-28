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

		cocktailGroup := cocktailsGroup.Group("/:cocktailId")
		{
			cocktailGroup.GET("", cocktailController.GetCocktail)
			cocktailGroup.DELETE(
				"",
				middleware.EnsureValidToken(scope.DeleteCocktail),
				cocktailController.DeleteCocktail,
			)

			cocktailGroup.PUT(
				"",
				middleware.EnsureValidToken(scope.UpdateCocktail),
				cocktailController.UpdateCocktail,
			)

			cocktailIngredientsGroup := cocktailGroup.Group("/ingredients")
			{
				cocktailIngredientGroup := cocktailIngredientsGroup.Group("/:ingredientId")
				{
					cocktailIngredientController := new(controllers.CocktailIngredientController)

					cocktailIngredientGroup.POST(
						"",
						middleware.EnsureValidToken(scope.CreateCocktailIngredient),
						cocktailIngredientController.CreateCocktailIngredient,
					)
					cocktailIngredientGroup.DELETE(
						"",
						middleware.EnsureValidToken(scope.DeleteCocktailIngredient),
						cocktailIngredientController.DeleteCocktailIngredient,
					)
					cocktailIngredientGroup.PUT(
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

		ingredientGroup := ingredientsGroup.Group("/:ingredientId")
		{
			ingredientGroup.PUT(
				"",
				middleware.EnsureValidToken(scope.UpdateIngredient),
				ingredientController.UpdateIngredient,
			)

			ingredientGroup.DELETE(
				"",
				middleware.EnsureValidToken(scope.DeleteIngredient),
				ingredientController.DeleteIngredient,
			)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
