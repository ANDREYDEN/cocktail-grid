package dtos

type IngredientDto struct {
	Name string `json:"name"`
}

type IngredientResultDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
