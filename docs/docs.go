// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/genre/add/": {
            "post": {
                "description": "Сreating a movie record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create"
                ],
                "summary": "CreateGenre",
                "parameters": [
                    {
                        "description": "Сreating a genre",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Genre"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/genre/delete/{id}": {
            "post": {
                "description": "Deleting data from a table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete"
                ],
                "summary": "DeleteGenres",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id genre",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/genre/update/{id}": {
            "post": {
                "description": "Updating values in the table genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update"
                ],
                "summary": "UpdateGenres",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id genre",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New values",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Genre"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/get": {
            "get": {
                "description": "Привествие",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hello"
                ],
                "summary": "Hello",
                "responses": {}
            }
        },
        "/get/Movie/{limit}": {
            "get": {
                "description": "Getting a list of movies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get"
                ],
                "summary": "GetListMovies",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Number of films",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/get/genre/{id}": {
            "get": {
                "description": "Getting genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get"
                ],
                "summary": "Get genre",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id genre",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/get/movie/{id}": {
            "get": {
                "description": "Getting the name of the movie and its description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get"
                ],
                "summary": "Get movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id movie",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/get/movieByGenres/{id}": {
            "get": {
                "description": "Getting movies with the selected genre",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get"
                ],
                "summary": "GetMovieByGenre",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id genre",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/movie/add/": {
            "post": {
                "description": "Сreating a movie record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create"
                ],
                "summary": "CreateMovie",
                "parameters": [
                    {
                        "description": "Title and description of the film",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Movie"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/movie/delete/{id}": {
            "post": {
                "description": "Deleting data from a table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete"
                ],
                "summary": "DeleteMovie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id movie",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/movie/update/{id}": {
            "post": {
                "description": "Updating values in the table movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update"
                ],
                "summary": "UpdateMovie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id movie",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New values",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Movie"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/movieGenres/add/": {
            "post": {
                "description": "Сreating a movie record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create"
                ],
                "summary": "CreateMovieGenres",
                "parameters": [
                    {
                        "description": "Linking genres to a film",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Movie_genre"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/movieGenres/delete/{id}": {
            "post": {
                "description": "Deleting data from a table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete"
                ],
                "summary": "DeleteMovieGenres",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id movieGenres",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/movieGenres/update/{id}": {
            "post": {
                "description": "Updating values in the table movieGenres",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update"
                ],
                "summary": "UpdateMovieGenres",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id movie_genre",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New values",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Movie_genre"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.Genre": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name_genre": {
                    "type": "string"
                }
            }
        },
        "model.Movie": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name_movie": {
                    "type": "string"
                }
            }
        },
        "model.Movie_genre": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "id_genre": {
                    "type": "string"
                },
                "id_movie": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5050",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
