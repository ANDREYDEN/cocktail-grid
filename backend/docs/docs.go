// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cocktails": {
            "get": {
                "description": "Retrieves all available cocktails",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cocktails"
                ],
                "summary": "Gets all cocktails",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.CocktailDto"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new cocktail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cocktails"
                ],
                "summary": "Creates a cocktail",
                "parameters": [
                    {
                        "description": "Cocktail object",
                        "name": "cocktail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CreateCocktailDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.CocktailDto"
                        }
                    }
                }
            }
        },
        "/cocktails/{cocktailId}/ingredients/{ingredientId}": {
            "post": {
                "description": "Adds an existing ingredient to the cocktail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cocktails"
                ],
                "summary": "Adds an ingredient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cocktail ID",
                        "name": "cocktailId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Ingredient ID",
                        "name": "ingredientId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Cocktail ingredient object",
                        "name": "cocktail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CocktailIngredientDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.CocktailIngredientResultDto"
                        }
                    }
                }
            }
        },
        "/ingredients": {
            "get": {
                "description": "Retrieves all available ingredients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredients"
                ],
                "summary": "Get all ingredients",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.IngredientResultDto"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new ingredient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredients"
                ],
                "summary": "Creates an ingredient",
                "parameters": [
                    {
                        "description": "Ingredient object",
                        "name": "cocktail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.IngredientDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.IngredientResultDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.CocktailDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "imageUrl": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.CocktailIngredientResultDto"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dtos.CocktailIngredientDto": {
            "type": "object",
            "properties": {
                "quantity": {
                    "type": "number"
                }
            }
        },
        "dtos.CocktailIngredientResultDto": {
            "type": "object",
            "properties": {
                "cocktailId": {
                    "type": "integer"
                },
                "ingredientId": {
                    "type": "integer"
                },
                "ingredientName": {
                    "type": "string"
                },
                "quantity": {
                    "type": "number"
                }
            }
        },
        "dtos.CreateCocktailDto": {
            "type": "object",
            "properties": {
                "imageUrl": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dtos.IngredientDto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "dtos.IngredientResultDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CocktailGrid API",
	Description:      "An API to get cocktails and ingredients",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
