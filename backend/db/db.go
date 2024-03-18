package db

import (
	"cocktail-grid/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	database = db
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Cocktail{}, &models.Ingredient{}, &models.CocktailIngredient{})
}

func GetDB() *gorm.DB {
	return database
}