package controllers

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/dtos"
	"cocktail-grid/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CocktailController struct{}

// List godoc
// @Summary Gets all cocktails
// @Schemes
// @Description Retrieves all available cocktails
// @Tags cocktails
// @Accept json
// @Produce json
// @Success 200 {object} dtos.CocktailDto
// @Router /cocktails [get]
func (cocktailController CocktailController) List(context *gin.Context) {
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
