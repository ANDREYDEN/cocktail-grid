package models

type CocktailIngredient struct {
	CocktailID   int `gorm:"primaryKey"`
	IngredientID int `gorm:"primaryKey"`
	Quantity     float32

	Ingredient Ingredient
	Cocktail   Cocktail
}
