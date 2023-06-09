package models

type cocktail struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}

var cocktails = []cocktail{
	{ID: 1, Title: "Mojito", ImageURL: "https://www.thecocktaildb.com/images/media/drink/3z6xdi1589574608.jpg"},
	{ID: 2, Title: "Margarita", ImageURL: "https://www.thecocktaildb.com/images/media/drink/5noda61589575158.jpg"},
	{ID: 3, Title: "Long Island Iced Tea", ImageURL: "https://www.thecocktaildb.com/images/media/drink/loezxn1504373874.jpg"},
}

func GetAllCocktails() []cocktail {
	return cocktails
}