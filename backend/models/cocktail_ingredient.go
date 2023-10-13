package models

type CocktailIngredient struct {
	CocktailID   uint `gorm:"primaryKey"`
	IngredientID uint `gorm:"primaryKey"`
	Quantity     float32

	Ingredient Ingredient
	Cocktail   Cocktail
}
