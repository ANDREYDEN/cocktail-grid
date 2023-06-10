package models

type CocktailIngredient struct {
	CocktailID   int `json:"cocktail_id"`
	IngredientID int `json:"ingredient_id"`
	Quantity     int `json:"quantity"`
}