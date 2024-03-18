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
                "parameters": [
                    {
                        "type": "string",
                        "description": "exclude ingredients",
                        "name": "compact",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/vms.CocktailVm"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing cocktail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cocktails"
                ],
                "summary": "Update a cocktail",
                "parameters": [
                    {
                        "description": "Cocktail object",
                        "name": "cocktail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.CocktailDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vms.CocktailVm"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/vms.CocktailVm"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                            "$ref": "#/definitions/vms.CocktailVm"
                        }
                    }
                }
            }
        },
        "/cocktails/{cocktailId}/ingredients/{ingredientId}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
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
                            "$ref": "#/definitions/vms.CocktailIngredientVm"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an ingredient from a cocktail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cocktails"
                ],
                "summary": "Deletes an ingredient from a cocktail",
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
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "object"
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/vms.IngredientVm"
                            }
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
                            "$ref": "#/definitions/vms.IngredientVm"
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
        "vms.CocktailIngredientVm": {
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
        "vms.CocktailVm": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "imageUrl": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "vms.DetailedCocktailVm": {
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
                        "$ref": "#/definitions/vms.CocktailIngredientVm"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "vms.IngredientVm": {
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
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
