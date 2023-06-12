package dtos

type CreateCocktailDto struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}

type CocktailDto struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	Ingredients []CocktailIngredientResultDto `json:"ingredients"`
}