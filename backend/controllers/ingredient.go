package controllers

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/dtos"
	"cocktail-grid/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IngredientController struct{}

// CreateIngredient godoc
// @Summary Creates an ingredient
// @Schemes
// @Description Creates a new ingredient
// @Tags ingredient
// @Accept json
// @Produce json
// @Param cocktail body dtos.IngredientDto true "Ingredient object"
// @Success 201 {object} dtos.IngredientResultDto
// @Router /ingredients [post]
func (cocktailController IngredientController) CreateIngredient(ctx *gin.Context) {
	var ingredientDto dtos.IngredientDto

	if err := ctx.BindJSON(&ingredientDto); err != nil {
		panic(err)
	}

	var ingredient = models.Ingredient{
		Name: ingredientDto.Name,
	}

	db := db.GetDB()
	db.Create(&ingredient)

	ingredientResultDto := dtos.IngredientResultDto{
		Name: ingredient.Name,
	}

	ctx.IndentedJSON(http.StatusCreated, ingredientResultDto)
}
