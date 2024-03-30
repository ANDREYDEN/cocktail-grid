package controllers

import (
	"cocktail-grid/backend/db"
	"cocktail-grid/backend/dtos"
	"cocktail-grid/backend/models"
	slice_utils "cocktail-grid/backend/utils"
	"cocktail-grid/backend/vms"
	"net/http"

	"github.com/gin-gonic/gin"
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
	db.Find(&ingredients)

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
