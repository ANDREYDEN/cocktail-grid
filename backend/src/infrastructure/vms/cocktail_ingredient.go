package vms

import "cocktail-grid/backend/src/domain/models"

type CocktailIngredientVm struct {
	CocktailID     uint    `json:"cocktailId"`
	IngredientID   uint    `json:"ingredientId"`
	IngredientName string  `json:"ingredientName"`
	Quantity       float32 `json:"quantity"`
}

func FromCocktailIngredientToVm(cocktailIngredient models.CocktailIngredient) CocktailIngredientVm {
	return CocktailIngredientVm{
		CocktailID:     cocktailIngredient.CocktailID,
		IngredientID:   cocktailIngredient.IngredientID,
		IngredientName: cocktailIngredient.Ingredient.Name,
		Quantity:       cocktailIngredient.Quantity,
	}
}
