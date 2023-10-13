package dtos

import (
	"cocktail-grid/backend/models"
)

type CocktailIngredientDto struct {
	Quantity float32 `json:"quantity"`
}

type CocktailIngredientResultDto struct {
	CocktailID     uint     `json:"cocktailId"`
	IngredientID   uint     `json:"ingredientId"`
	IngredientName string  `json:"ingredie ntName"`
	Quantity       float32 `json:"quantity"`
}

func FromCocktailIngredientToDto(cocktailIngredient models.CocktailIngredient) CocktailIngredientResultDto {
	return CocktailIngredientResultDto {
		CocktailID:     cocktailIngredient.CocktailID,
		IngredientID:   cocktailIngredient.IngredientID,
		IngredientName: cocktailIngredient.Ingredient.Name,
		Quantity:       cocktailIngredient.Quantity,
	}
}
