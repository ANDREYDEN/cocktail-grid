package scope

import "fmt"

type Scope struct {
	Verb string
	Noun string
}

func (s Scope) ToString() string {
	return fmt.Sprintf("%s:%s", s.Verb, s.Noun)
}

func (s Scope) IsEmpty() bool {
	return s.Verb == "" && s.Noun == ""
}

var CreateCocktail Scope
var UpdateCocktail Scope
var DeleteCocktail Scope
var CreateIngredient Scope
var DeleteIngredient Scope
var CreateCocktailIngredient Scope
var DeleteCocktailIngredient Scope
var UpdateCocktailIngredient Scope

func InitScopes() {
	create := "create"
	update := "update"
	delete := "delete"

	cocktail := "cocktail"
	ingredient := "ingredient"
	cocktailIngredient := "cocktail_ingredient"

	CreateCocktail = Scope{Verb: create, Noun: cocktail}
	UpdateCocktail = Scope{Verb: update, Noun: cocktail}
	DeleteCocktail = Scope{Verb: delete, Noun: cocktail}
	CreateIngredient = Scope{Verb: create, Noun: ingredient}
	DeleteIngredient = Scope{Verb: delete, Noun: ingredient}
	CreateCocktailIngredient = Scope{Verb: create, Noun: cocktailIngredient}
	DeleteCocktailIngredient = Scope{Verb: delete, Noun: cocktailIngredient}
	UpdateCocktailIngredient = Scope{Verb: update, Noun: cocktailIngredient}
}
