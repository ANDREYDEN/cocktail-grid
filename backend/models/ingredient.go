package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name string 

	CocktailIngredients []CocktailIngredient
}