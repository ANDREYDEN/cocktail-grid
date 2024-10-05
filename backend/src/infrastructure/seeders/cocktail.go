package seeders

import (
	"cocktail-grid/backend/src/domain/models"

	"gorm.io/gorm"
)

func SeedCocktail(db *gorm.DB, cocktail *models.Cocktail) error {
	return db.Create(cocktail).Error
}
