package controllers

import (
	"cocktail-grid/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CocktailController struct {}

// getCocktails godoc
// @Summary Gets all cocktails
// @Schemes
// @Description Retrieves all available cocktails
// @Tags cocktails
// @Accept json
// @Produce json
// @Success 200 {object} models.Cocktail
// @Router /cocktails [get]
func (cocktailController CocktailController) GetCocktails(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.GetAllCocktails())
}