package controllers

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/dtos"
	"cocktail-grid/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IngredientController struct{}

// GetAllIngredients godoc
// @Summary Get all ingredients
// @Schemes
// @Description Retrieves all available ingredients
// @Tags Ingredients
// @Accept json
// @Produce json
// @Success 200 {object} dtos.IngredientResultDto
// @Router /ingredients [get]
func (cocktailController IngredientController) GetAllIngredients(ctx *gin.Context) {
	var ingredients []models.Ingredient

	db := db.GetDB()
	db.Find(&ingredients)

	ingredientResultDtos := []dtos.IngredientResultDto{}
	for _, ingredient := range ingredients {
		ingredientResultDto := dtos.IngredientResultDto{
			ID:   ingredient.ID,
			Name: ingredient.Name,
		}
		ingredientResultDtos = append(ingredientResultDtos, ingredientResultDto)
	}

	ctx.IndentedJSON(http.StatusOK, ingredientResultDtos)
}

// CreateIngredient godoc
// @Summary Creates an ingredient
// @Schemes
// @Description Creates a new ingredient
// @Tags Ingredients
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
