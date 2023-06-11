package controllers

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/dtos"
	"cocktail-grid/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CocktailController struct{}

// GetAll godoc
// @Summary Gets all cocktails
// @Schemes
// @Description Retrieves all available cocktails
// @Tags cocktails
// @Accept json
// @Produce json
// @Success 200 {object} dtos.CocktailDto
// @Router /cocktails [get]
func (cocktailController CocktailController) GetAll(context *gin.Context) {
	db := db.GetDB()
	var cocktails []models.Cocktail
	db.Find(&cocktails)

	cocktailDtos := []dtos.CocktailDto{}
	for _, cocktail := range cocktails {
		cocktailDto := dtos.CocktailDto{
			Title:    cocktail.Title,
			ImageURL: cocktail.ImageURL,
		}
		cocktailDtos = append(cocktailDtos, cocktailDto)
	}

	context.IndentedJSON(http.StatusOK, cocktailDtos)
}

// Create godoc
// @Summary Creates a cocktail
// @Schemes
// @Description Creates a new cocktail
// @Tags cocktails
// @Accept json
// @Produce json
// @Param cocktail body dtos.CocktailDto true "Cocktail object"
// @Success 201 {object} dtos.CocktailDto
// @Router /cocktails [post]
func (cocktailController CocktailController) Create(context *gin.Context) {
	var cocktailDto dtos.CocktailDto

	if err := context.BindJSON(&cocktailDto); err != nil {
		return
	}

	cocktail := models.Cocktail{
		Title:    cocktailDto.Title,
		ImageURL: cocktailDto.ImageURL,
	}

	db := db.GetDB()
	db.Create(&cocktail)

	context.IndentedJSON(http.StatusCreated, cocktailDto)
}
