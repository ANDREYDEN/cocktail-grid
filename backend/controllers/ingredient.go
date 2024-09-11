package controllers

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/dtos"
	"cocktail-grid/backend/models"
	slice_utils "cocktail-grid/backend/utils"
	"cocktail-grid/backend/vms"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IngredientController struct{}

// GetAllIngredients godoc
//
//	@Summary	Get all ingredients
//	@Schemes
//	@Description	Retrieves all available ingredients
//	@Tags			Ingredients
//	@ID				Get_Ingredients
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]vms.IngredientVm
//	@Router			/ingredients [get]
func (cocktailController IngredientController) GetAllIngredients(ctx *gin.Context) {
	var ingredients []models.Ingredient

	db := db.GetDB()
	db.Order("LOWER(name)").Find(&ingredients)

	ingredientVms := slice_utils.Map(ingredients, vms.FromIngredientToVm)

	ctx.IndentedJSON(http.StatusOK, ingredientVms)
}

// CreateIngredient godoc
//
//	@Summary	Creates an ingredient
//	@Schemes
//	@Description	Creates a new ingredient
//	@Tags			Ingredients
//	@ID				Create_Ingredient
//	@Accept			json
//	@Produce		json
//	@Param			cocktail	body		dtos.IngredientDto	true	"Ingredient object"
//	@Success		201			{object}	vms.IngredientVm
//	@Router			/ingredients [post]
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

	ctx.IndentedJSON(http.StatusCreated, vms.FromIngredientToVm(ingredient))
}

// UpdateIngredient godoc
//
//	@Summary	Updates an ingredient
//	@Schemes
//	@Description	Updates an ingredient
//	@Tags			Ingredients
//	@ID				Update_Ingredient
//	@Accept			json
//	@Produce		json
//	@Param			cocktail		body		dtos.IngredientDto	true	"Ingredient object"
//	@Param			ingredientId	path		int					true	"Ingredient ID"
//	@Success		200				{object}	interface{}
//	@Router			/ingredients/{ingredientId} [put]
//	@Security		BearerAuth
func (ingredientController IngredientController) UpdateIngredient(ctx *gin.Context) {
	type UpdateIngredientPathParams struct {
		IngredientId uint `uri:"ingredientId" binding:"required"`
	}

	pathParams, ok := GetPathParams[UpdateIngredientPathParams](ctx)
	if !ok {
		return
	}

	var ingredientDto dtos.IngredientDto
	if err := ctx.BindJSON(&ingredientDto); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	var ingredient models.Ingredient
	result := db.First(&ingredient, pathParams.IngredientId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No ingredient found with id %d", ingredient.ID)})
		return
	}

	ingredient.Name = ingredientDto.Name

	db.Save(&ingredient)

	ctx.IndentedJSON(http.StatusOK, vms.FromIngredientToVm(ingredient))
}

// DeleteIngredient godoc
//
//	@Summary	Deletes an ingredient
//	@Schemes
//	@Description	Deletes an ingredient
//	@Tags			Ingredients
//	@ID				Delete_Ingredient
//	@Accept			json
//	@Produce		json
//	@Param			ingredientId	path		int	true	"Ingredient ID"
//	@Success		200				{object}	interface{}
//	@Router			/ingredients/{ingredientId} [delete]
//	@Security		BearerAuth
func (ingredientController IngredientController) DeleteIngredient(ctx *gin.Context) {
	type DeleteIngredientPathParams struct {
		IngredientId uint `uri:"ingredientId" binding:"required"`
	}

	pathParams, ok := GetPathParams[DeleteIngredientPathParams](ctx)
	if !ok {
		return
	}

	err := db.GetDB().First(&models.Ingredient{}, pathParams.IngredientId).Delete(&models.Ingredient{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ingredient (%d) was not found", pathParams.IngredientId)})
			return
		}
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, gin.H{})
}
