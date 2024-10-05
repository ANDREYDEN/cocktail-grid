package controllers

import (
	"cocktail-grid/backend/src/domain/models"
	"cocktail-grid/backend/src/infrastructure/db"
	"cocktail-grid/backend/src/infrastructure/dtos"
	"cocktail-grid/backend/src/infrastructure/vms"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CocktailIngredientController struct{}

type CRUDCocktailIngredientPathParams struct {
	CocktailID   uint `uri:"cocktailId" binding:"required"`
	IngredientID uint `uri:"ingredientId" binding:"required"`
}

// CreateCocktailIngredient godoc
//
//	@Summary	Adds an ingredient
//	@Schemes
//	@Description	Adds an existing ingredient to the cocktail
//	@Tags			Cocktails
//	@ID				Create_CocktailIngredient
//	@Accept			json
//	@Produce		json
//	@Param			cocktailId			path		int							true	"Cocktail ID"
//	@Param			ingredientId		path		int							true	"Ingredient ID"
//	@Param			cocktailIngredient	body		dtos.CocktailIngredientDto	true	"Cocktail ingredient object"
//	@Success		201					{object}	vms.CocktailIngredientVm
//	@Failure		422					{object}	interface{}
//	@Router			/cocktails/{cocktailId}/ingredients/{ingredientId} [post]
//	@Security		BearerAuth
func (cocktailController CocktailIngredientController) CreateCocktailIngredient(ctx *gin.Context) {
	pathParams, ok := GetPathParams[CRUDCocktailIngredientPathParams](ctx)
	if !ok {
		return
	}

	var cocktailIngredientDto dtos.CocktailIngredientDto

	if err := ctx.BindJSON(&cocktailIngredientDto); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	var cocktailIngredient = models.CocktailIngredient{
		CocktailID:   pathParams.CocktailID,
		IngredientID: pathParams.IngredientID,
		Quantity:     cocktailIngredientDto.Quantity,
	}

	db := db.GetDB()
	result := db.Create(&cocktailIngredient)
	if result.Error != nil {
		// TODO: make this check better when SQLITE is no longer used
		if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
			ctx.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{
					"error": fmt.Sprintf(
						"Ingredient (id: %d) already exists as part of this cocktail (id: %d)",
						cocktailIngredient.IngredientID,
						cocktailIngredient.CocktailID,
					),
				},
			)
			return
		}
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{"error": result.Error.Error()},
		)
		return
	}
	ctx.IndentedJSON(http.StatusCreated, vms.FromCocktailIngredientToVm(cocktailIngredient))
}

// DeleteCocktailIngredient godoc
//
//	@Summary	Deletes an ingredient from a cocktail
//	@Schemes
//	@Description	Deletes an ingredient from a cocktail
//	@Tags			Cocktails
//	@ID				Delete_CocktailIngredient
//	@Accept			json
//	@Produce		json
//	@Param			cocktailId		path		int	true	"Cocktail ID"
//	@Param			ingredientId	path		int	true	"Ingredient ID"
//	@Success		201				{object}	interface{}
//	@Router			/cocktails/{cocktailId}/ingredients/{ingredientId} [delete]
//	@Security		BearerAuth
func (cocktailController CocktailIngredientController) DeleteCocktailIngredient(ctx *gin.Context) {
	pathParams, ok := GetPathParams[CRUDCocktailIngredientPathParams](ctx)
	if !ok {
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

// UpdateCocktailIngredient godoc
//
//	@Summary	Updates a cocktail ingredient
//	@Schemes
//	@Description	Updates a cocktail ingredient
//	@Tags			Cocktails
//	@ID				Update_CocktailIngredient
//	@Accept			json
//	@Produce		json
//	@Param			cocktailId			path		int							true	"Cocktail ID"
//	@Param			ingredientId		path		int							true	"Ingredient ID"
//	@Param			cocktailIngredient	body		dtos.CocktailIngredientDto	true	"Cocktail ingredient object"
//	@Success		200					{object}	interface{}
//	@Router			/cocktails/{cocktailId}/ingredients/{ingredientId} [put]
//	@Security		BearerAuth
func (cocktailController CocktailIngredientController) UpdateCocktailIngredient(ctx *gin.Context) {
	pathParams, ok := GetPathParams[CRUDCocktailIngredientPathParams](ctx)
	if !ok {
		return
	}

	var cocktailIngredientDto dtos.CocktailIngredientDto
	if err := ctx.BindJSON(&cocktailIngredientDto); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var cocktailIngredient = models.CocktailIngredient{
		CocktailID:   pathParams.CocktailID,
		IngredientID: pathParams.IngredientID,
	}

	db := db.GetDB()

	err := db.First(&cocktailIngredient).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Cocktail ingredient (%d, %d) was not found", pathParams.CocktailID, pathParams.IngredientID)})
			return
		}
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	cocktailIngredient.Quantity = cocktailIngredientDto.Quantity

	db.Save(&cocktailIngredient)

	ctx.IndentedJSON(http.StatusOK, vms.FromCocktailIngredientToVm(cocktailIngredient))
}
