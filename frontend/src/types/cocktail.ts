export type CocktailDto = {
    id: number;
    title: string;
    imageUrl: string;
    ingredients: CocktailIngredientResultDto[];
}

type CocktailIngredientResultDto = {
    cocktailId: number
	ingredientId: number
	ingredientName: string
	quantity: number
}