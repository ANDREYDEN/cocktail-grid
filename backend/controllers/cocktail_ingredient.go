package controllers

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/dtos"
	"cocktail-grid/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CocktailIngredientController struct{}

// CreateCocktailIngredient godoc
// @Summary Adds an ingredient
// @Schemes
// @Description Adds an existing ingredient to the cocktail
// @Tags Cocktails
// @Accept json
// @Produce json
// @Param        cocktailId   path      int  true  "Cocktail ID"
// @Param        ingredientId   path      int  true  "Ingredient ID"
// @Param cocktail body dtos.CocktailIngredientDto true "Cocktail ingredient object"
// @Success 201 {object} dtos.CocktailIngredientResultDto
// @Router /cocktails/{cocktailId}/ingredients/{ingredientId} [post]
func (cocktailController CocktailIngredientController) CreateCocktailIngredient(ctx *gin.Context) {
	type CreateCocktailIngredientPathParams struct {
		CocktailID   int `uri:"cocktailId" binding:"required"`
		IngredientID int `uri:"ingredientId" binding:"required"`
	}

	var pathParams CreateCocktailIngredientPathParams

	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	var cocktailIngredientDto dtos.CocktailIngredientDto

	if err := ctx.BindJSON(&cocktailIngredientDto); err != nil {
		panic(err)
	}

	var cocktailIngredient = models.CocktailIngredient{
		CocktailID:   pathParams.CocktailID,
		IngredientID: pathParams.IngredientID,
		Quantity:     cocktailIngredientDto.Quantity,
	}

	db := db.GetDB()
	db.Create(&cocktailIngredient)

	cocktailIngredientResultDto := dtos.CocktailIngredientResultDto{
		CocktailID:     cocktailIngredient.CocktailID,
		IngredientID:   cocktailIngredient.IngredientID,
		IngredientName: cocktailIngredient.Ingredient.Name,
		Quantity:       cocktailIngredient.Quantity,
	}

	ctx.IndentedJSON(http.StatusCreated, cocktailIngredientResultDto)
}
