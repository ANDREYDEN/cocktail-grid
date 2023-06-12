package dtos

type IngredientDto struct {
	Name string `json:"name"`
}

type IngredientResultDto struct {
	ID   int   `json:"id"`
	Name string `json:"name"`
}
