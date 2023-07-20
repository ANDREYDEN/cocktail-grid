package dtos

import (
	"cocktail-grid/backend/models"
)

type CreateCocktailDto struct {
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
}

type CocktailDto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
	Ingredients []CocktailIngredientResultDto `json:"ingredients"`
}

// Maps a Cocktail model to a CocktailDto. 
// The cocktail should have its CocktailIngredients preloaded.
func FromCocktailToDto(cocktail models.Cocktail) CocktailDto {
	ingredientResultDtos := []CocktailIngredientResultDto{}
	for _, cocktailIngredient := range cocktail.CocktailIngredients {
		cocktailIngredientResultDto := FromCocktailIngredientToDto(cocktailIngredient)
		ingredientResultDtos = append(ingredientResultDtos, cocktailIngredientResultDto)
	}

	return CocktailDto{
		ID:	   cocktail.ID,
		Title:    cocktail.Title,
		ImageURL: cocktail.ImageURL,
		Ingredients: ingredientResultDtos,
	}
}