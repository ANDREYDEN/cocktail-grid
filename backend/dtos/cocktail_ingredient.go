package dtos

type CocktailIngredientDto struct {
	Quantity float32 `json:"quantity"`
}

type CocktailIngredientResultDto struct {
	CocktailID     int     `json:"cocktailId"`
	IngredientID   int     `json:"ingredientId"`
	IngredientName string  `json:"ingredientName"`
	Quantity       float32 `json:"quantity"`
}
