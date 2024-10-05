package vms

import "cocktail-grid/backend/src/domain/models"

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
