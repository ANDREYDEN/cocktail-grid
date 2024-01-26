package controllers

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/dtos"
	"cocktail-grid/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CocktailController struct{}

// GetAllCocktails godoc
//	@Summary	Gets all cocktails
//	@Schemes
//	@Description	Retrieves all available cocktails
//	@Tags			Cocktails
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dtos.CocktailDto
//	@Router			/cocktails [get]
func (cocktailController CocktailController) GetAllCocktails(ctx *gin.Context) {
	db := db.GetDB()
	var cocktails []models.Cocktail
	db.Preload("CocktailIngredients.Ingredient").Find(&cocktails)

	cocktailDtos := []dtos.CocktailDto{}
	for _, cocktail := range cocktails {
		cocktailDto := dtos.FromCocktailToDto(cocktail)
		cocktailDtos = append(cocktailDtos, cocktailDto)
	}

	ctx.IndentedJSON(http.StatusOK, cocktailDtos)
}

// CreateCocktail godoc
//	@Summary	Creates a cocktail
//	@Schemes
//	@Description	Creates a new cocktail
//	@Tags			Cocktails
//	@Accept			json
//	@Produce		json
//	@Param			cocktail	body		dtos.CreateCocktailDto	true	"Cocktail object"
//	@Success		201			{object}	dtos.CocktailDto
//	@Router			/cocktails [post]
func (cocktailController CocktailController) CreateCocktail(ctx *gin.Context) {
	var cocktailDto dtos.CocktailDto

	if err := ctx.BindJSON(&cocktailDto); err != nil {
		return
	}

	cocktail := models.Cocktail{
		Title:    cocktailDto.Title,
		ImageURL: cocktailDto.ImageURL,
	}

	db := db.GetDB()
	db.Create(&cocktail)

	ctx.IndentedJSON(http.StatusCreated, cocktailDto)
}
