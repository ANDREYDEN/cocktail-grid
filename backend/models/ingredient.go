package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	ID   int    
	Name string 

	CocktailIngredients []CocktailIngredient
}