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

type CocktailController struct{}

// GetAllCocktails godoc
//	@Summary	Gets all cocktails
//	@Schemes
//	@Description	Retrieves all available cocktails
//	@Tags			Cocktails
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]vms.DetailedCocktailVm
//	@Router			/cocktails [get]
func (cocktailController CocktailController) GetAllCocktails(ctx *gin.Context) {
	db := db.GetDB()
	var cocktails []models.Cocktail
	db.Preload("CocktailIngredients.Ingredient").Find(&cocktails)

	cocktailVms := []vms.DetailedCocktailVm{}
	for _, cocktail := range cocktails {
		cocktailDto := vms.FromCocktailToDetailedVm(cocktail)
		cocktailVms = append(cocktailVms, cocktailDto)
	}

	ctx.IndentedJSON(http.StatusOK, cocktailVms)
}

// CreateCocktail godoc
//	@Summary	Creates a cocktail
//	@Schemes
//	@Description	Creates a new cocktail
//	@Tags			Cocktails
//	@Accept			json
//	@Produce		json
//	@Param			cocktail	body		dtos.CreateCocktailDto	true	"Cocktail object"
//	@Success		201			{object}	vms.CocktailVm
//	@Router			/cocktails [post]
func (cocktailController CocktailController) CreateCocktail(ctx *gin.Context) {
	var cocktailDto dtos.CocktailDto

	if err := ctx.BindJSON(&cocktailDto); err != nil {
		return
	}

	cocktail := dtos.FromDtoToCocktail(cocktailDto) 

	db := db.GetDB()
	db.Create(&cocktail)

	ctx.IndentedJSON(http.StatusCreated, vms.FromCocktailToVm(cocktail))
}

// UpdateCocktail godoc
//	@Summary	Update a cocktail
//	@Schemes
//	@Description	Update an existing cocktail
//	@Tags			Cocktails
//	@Accept			json
//	@Produce		json
//	@Param			cocktail	body		dtos.CocktailDto	true	"Cocktail object"
//	@Success		200			{object}    vms.CocktailVm	
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