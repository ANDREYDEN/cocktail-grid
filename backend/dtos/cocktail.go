package dtos

import "cocktail-grid/backend/models"

type CreateCocktailDto struct {
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
}

type CocktailDto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"imageUrl"`
}

func FromDtoToCocktail(cocktailDto CocktailDto) models.Cocktail {
	return models.Cocktail {
		Title: cocktailDto.Title,
		ImageURL: cocktailDto.ImageURL,
	} 
}