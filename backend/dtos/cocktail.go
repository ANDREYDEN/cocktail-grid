package dtos

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