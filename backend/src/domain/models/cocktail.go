package models

import "gorm.io/gorm"

type Cocktail struct {
	gorm.Model
	Title    string
	ImageURL string
	
	CocktailIngredients []CocktailIngredient
}