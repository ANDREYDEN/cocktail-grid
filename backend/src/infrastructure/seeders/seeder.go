package seeders

import (
	"cocktail-grid/backend/src/domain/models"
	"cocktail-grid/backend/src/infrastructure/db"
	"log"
)

func Seed() {
	db.Init()
	db := db.GetDB()

	log.Printf("Running Cocktail Seeder...")
	margarita := models.Cocktail{
		Title: "Margarita",
	}
	mojito := models.Cocktail{
		Title: "Mojito",
	}
	SeedCocktail(db, &margarita)
	SeedCocktail(db, &mojito)

	log.Printf("Running Ingredient Seeder...")
	tequila := models.Ingredient{
		Name: "Silver Tequila",
	}
	limeJuice := models.Ingredient{
		Name: "Lime Juice",
	}
	tripleSec := models.Ingredient{
		Name: "Triple Sec",
	}
	whiteRum := models.Ingredient{
		Name: "White Rum",
	}
	simpleSyrup := models.Ingredient{
		Name: "Simple Syrup",
	}
	clubSoda := models.Ingredient{
		Name: "Club Soda",
	}
	SeedIngredient(db, &tequila)
	SeedIngredient(db, &limeJuice)
	SeedIngredient(db, &tripleSec)
	SeedIngredient(db, &whiteRum)
	SeedIngredient(db, &simpleSyrup)
	SeedIngredient(db, &clubSoda)

	log.Printf("Running Cocktail Ingredient Seeder...")
	SeedCocktailIngredient(db, margarita.ID, tequila.ID, 1.5)
	SeedCocktailIngredient(db, margarita.ID, limeJuice.ID, 0.75)
	SeedCocktailIngredient(db, margarita.ID, tripleSec.ID, 1)

	SeedCocktailIngredient(db, mojito.ID, whiteRum.ID, 2)
	SeedCocktailIngredient(db, mojito.ID, limeJuice.ID, 1)
	SeedCocktailIngredient(db, mojito.ID, simpleSyrup.ID, 0.5)
	SeedCocktailIngredient(db, mojito.ID, clubSoda.ID, -1)
}
