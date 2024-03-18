package controllers

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/dtos"
	"cocktail-grid/backend/models"
	"cocktail-grid/backend/vms"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CocktailIngredientController struct{}

// CreateCocktailIngredient godoc
//
//	@Summary	Adds an ingredient
//	@Schemes
//	@Description	Adds an existing ingredient to the cocktail
//	@Tags			Cocktails
//	@Accept			json
//	@Produce		json
//	@Param			cocktailId		path		int							true	"Cocktail ID"
//	@Param			ingredientId	path		int							true	"Ingredient ID"
//	@Param			cocktail		body		dtos.CocktailIngredientDto	true	"Cocktail ingredient object"
//	@Success		201				{object}	vms.CocktailIngredientVm
//	@Router			/cocktails/{cocktailId}/ingredients/{ingredientId} [post]
//	@Security		BearerAuth
func (cocktailController CocktailIngredientController) CreateCocktailIngredient(ctx *gin.Context) {
	type CreateCocktailIngredientPathParams struct {
		CocktailID   uint `uri:"cocktailId" binding:"required"`
		IngredientID uint `uri:"ingredientId" binding:"required"`
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

	ctx.IndentedJSON(http.StatusCreated, vms.FromCocktailIngredientToVm(cocktailIngredient))
}

// DeleteCocktailIngredient godoc
//
//	@Summary	Deletes an ingredient from a cocktail
//	@Schemes
//	@Description	Deletes an ingredient from a cocktail
//	@Tags			Cocktails
//	@Accept			json
//	@Produce		json
//	@Param			cocktailId		path		int	true	"Cocktail ID"
//	@Param			ingredientId	path		int	true	"Ingredient ID"
//	@Success		204				{object}	interface{}
//	@Router			/cocktails/{cocktailId}/ingredients/{ingredientId} [delete]
func (cocktailController CocktailIngredientController) DeleteCocktailIngredient(ctx *gin.Context) {
	type CreateCocktailIngredientPathParams struct {
		CocktailID   uint `uri:"cocktailId" binding:"required"`
		IngredientID uint `uri:"ingredientId" binding:"required"`
	}

	var pathParams CreateCocktailIngredientPathParams
	if err := ctx.ShouldBindUri(&pathParams); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var cocktailIngredient = models.CocktailIngredient{
		CocktailID:   pathParams.CocktailID,
		IngredientID: pathParams.IngredientID,
	}

	db := db.GetDB()

	err := db.First(&cocktailIngredient).Delete(&models.CocktailIngredient{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Cocktail ingredient (%d, %d) was not found", pathParams.CocktailID, pathParams.IngredientID)})
			return
		}
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, gin.H{})
}
