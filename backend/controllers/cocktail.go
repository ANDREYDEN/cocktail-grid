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

type CocktailController struct{}

// GetAllCocktails godoc
//
//	@Summary	Gets all cocktails
//	@Schemes
//	@Description	Retrieves all available cocktails
//	@Tags			Cocktails
//	@ID				Get_Cocktails
//	@Accept			json
//	@Produce		json
//	@Param			compact	query		string	false	"exclude ingredients"
//	@Success		200		{object}	[]vms.DetailedCocktailVm
//	@Router			/cocktails [get]
func (cocktailController CocktailController) GetAllCocktails(ctx *gin.Context) {
	db := db.GetDB()
	var cocktails []models.Cocktail
	compact := ctx.Query("compact")
	query := db
	if compact == "" {
		query = db.Preload("CocktailIngredients.Ingredient")
	}
	query.Find(&cocktails)

	if compact == "" {
		cocktailVms := slice_utils.Map(cocktails, vms.FromCocktailToDetailedVm)
		ctx.IndentedJSON(http.StatusOK, cocktailVms)
	} else {
		cocktailVms := slice_utils.Map(cocktails, vms.FromCocktailToVm)
		ctx.IndentedJSON(http.StatusOK, cocktailVms)
	}
}

// CreateCocktail godoc
//
//	@Summary	Creates a cocktail
//	@Schemes
//	@Description	Creates a new cocktail
//	@Tags			Cocktails
//	@ID				Create_Cocktail
//	@Accept			json
//	@Produce		json
//	@Param			cocktail	body		dtos.CreateCocktailDto	true	"Cocktail object"
//	@Success		201			{object}	vms.CocktailVm
//	@Router			/cocktails [post]
//	@Security		ApiKeyAuth
func (cocktailController CocktailController) CreateCocktail(ctx *gin.Context) {
	var cocktailDto dtos.CocktailDto

	if err := ctx.BindJSON(&cocktailDto); err != nil {
		return
	}

	if cocktailDto.Title == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Title is required"})
		return
	}

	cocktail := dtos.FromDtoToCocktail(cocktailDto)

	db := db.GetDB()
	db.Create(&cocktail)

	ctx.IndentedJSON(http.StatusCreated, vms.FromCocktailToVm(cocktail))
}

// UpdateCocktail godoc
//
//	@Summary	Update a cocktail
//	@Schemes
//	@Description	Update an existing cocktail
//	@Tags			Cocktails
//	@Accept			json
//	@Produce		json
//	@Param			cocktail	body		dtos.CocktailDto	true	"Cocktail object"
//	@Success		200			{object}	vms.CocktailVm
//	@Success		404			{object}	vms.CocktailVm
//	@Router			/cocktails [put]
func (cocktailController CocktailController) UpdateCocktail(ctx *gin.Context) {
	var cocktailDto dtos.CocktailDto
	if err := ctx.BindJSON(&cocktailDto); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	db := db.GetDB()
	var cocktail models.Cocktail
	result := db.First(&cocktail, cocktailDto.ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("No cocktail found with id %d", cocktail.ID)})
		return
	}

	cocktail.Title = cocktailDto.Title
	cocktail.ImageURL = cocktailDto.ImageURL

	db.Save(&cocktail)

	ctx.IndentedJSON(http.StatusOK, vms.FromCocktailToVm(cocktail))
}

// DeleteCocktail godoc
//
//	@Summary	Deletes a cocktail
//	@Schemes
//	@Description	Deletes a cocktail
//	@Tags			Cocktails
//	@ID				Delete_Cocktail
//	@Accept			json
//	@Produce		json
//	@Param			cocktailId	path		int	true	"Cocktail ID"
//	@Success		200			{object}	interface{}
//	@Router			/cocktails/{cocktailId} [delete]
//	@Security		BearerAuth
func (cocktailController CocktailController) DeleteCocktail(ctx *gin.Context) {
	type DeleteCocktailPathParams struct {
		CocktailId uint `uri:"cocktailId" binding:"required"`
	}

	pathParams, ok := GetPathParams[DeleteCocktailPathParams](ctx)
	if !ok {
		return
	}

	db := db.GetDB()

	err := db.First(&models.Cocktail{}, pathParams.CocktailId).Delete(&models.Cocktail{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Cocktail (%d) was not found", pathParams.CocktailId)})
			return
		}
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusNoContent, gin.H{})
}
