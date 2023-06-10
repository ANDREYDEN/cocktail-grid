package controllers

import (
	"cocktail-grid/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CocktailController struct {}

// getCocktails godoc
// @Summary retrieves cocktails
// @Schemes
// @Description retrieves all available cocktails
// @Tags cocktails
// @Accept json
// @Produce json
// @Success 200 {object} models.cocktail
// @Router /cocktails [get]
func (cocktailController CocktailController) GetCocktails(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, models.GetAllCocktails())
}