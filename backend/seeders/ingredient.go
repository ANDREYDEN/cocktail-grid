package seeders

import (
	"cocktail-grid/backend/models"

	"gorm.io/gorm"
)

func SeedIngredient(db *gorm.DB, ingredient *models.Ingredient) error {
	return db.Create(ingredient).Error
}