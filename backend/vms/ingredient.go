package vms

import "cocktail-grid/backend/models"

type IngredientVm struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func FromIngredientToVm(ingredient models.Ingredient) IngredientVm {
	return IngredientVm{
		ID:   ingredient.ID,
		Name: ingredient.Name,
	}
}
