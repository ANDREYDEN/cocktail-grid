package seeders

import (
	"cocktail-grid/backend/src/domain/models"

	"gorm.io/gorm"
)

func SeedCocktailIngredient(db *gorm.DB, cocktailId uint, ingredientId uint, quantity float32) {
	cocktailIngredient := models.CocktailIngredient{
		CocktailID:   cocktailId,
		IngredientID: ingredientId,
		Quantity:     quantity,
	}
	db.Create(&cocktailIngredient)
}
