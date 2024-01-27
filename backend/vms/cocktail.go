package vms

import "cocktail-grid/backend/models"

type CocktailVm struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
}

type DetailedCocktailVm struct {
	CocktailVm
	Ingredients []CocktailIngredientVm `json:"ingredients"`
}

// Maps a Cocktail model to a CocktailVm. 
// The cocktail should have its CocktailIngredients preloaded.
func FromCocktailToVm(cocktail models.Cocktail) CocktailVm {
	return CocktailVm{
		ID:	   cocktail.ID,
		Title:    cocktail.Title,
		ImageURL: cocktail.ImageURL,
	}
}


// Maps a Cocktail model to a CocktailVm. 
// The cocktail should have its CocktailIngredients preloaded.
func FromCocktailToDetailedVm(cocktail models.Cocktail) DetailedCocktailVm {
	ingredientVms := []CocktailIngredientVm{}
	for _, cocktailIngredient := range cocktail.CocktailIngredients {
		cocktailIngredientVm := FromCocktailIngredientToVm(cocktailIngredient)
		ingredientVms = append(ingredientVms, cocktailIngredientVm)
	}

	return DetailedCocktailVm{
	CocktailVm: FromCocktailToVm(cocktail),
		Ingredients: ingredientVms,
	}
}