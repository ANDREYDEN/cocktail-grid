package models

import "gorm.io/gorm"

type Cocktail struct {
	gorm.Model
	Title    string
	ImageURL string
	
	Ingredients []Ingredient `gorm:"many2many:cocktail_ingredients;"`
}