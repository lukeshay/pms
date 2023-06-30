// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/": {
            "get": {
                "description": "get jwt claims",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get jwt claims",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthGetResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Create user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthPostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthPostResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/auth/sign-in/": {
            "post": {
                "description": "sign in to user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Sign in to user",
                "parameters": [
                    {
                        "description": "Sign in to user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthPostSignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.AuthPostResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/books/": {
            "get": {
                "description": "get books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "List books",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BooksListResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "post book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Create book",
                "parameters": [
                    {
                        "description": "Create book",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.BooksPostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BooksGetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/books/{id}/": {
            "get": {
                "description": "get book by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Get book by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BooksGetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "put book by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Update book by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update book",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.BooksPostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BooksGetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete book by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Delete book by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.BooksDeleteResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}/": {
            "get": {
                "description": "get user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.UsersGetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "put user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UsersPostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.UsersGetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.Claims": {
            "type": "object",
            "required": [
                "createdAt",
                "createdBy",
                "email",
                "emailVerified",
                "firstName",
                "id",
                "lastName",
                "updatedAt",
                "updatedBy"
            ],
            "properties": {
                "aud": {
                    "description": "the ` + "`" + `aud` + "`" + ` (Audience) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.3",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "createdAt": {
                    "description": "When the record was created",
                    "type": "string"
                },
                "createdBy": {
                    "description": "Who created the record",
                    "type": "string"
                },
                "deletedAt": {
                    "description": "When the record was deleted",
                    "type": "string"
                },
                "deletedBy": {
                    "description": "Who deleted the record",
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "emailVerified": {
                    "type": "boolean"
                },
                "exp": {
                    "description": "the ` + "`" + `exp` + "`" + ` (Expiration Time) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4",
                    "allOf": [
                        {
                            "$ref": "#/definitions/jwt.NumericDate"
                        }
                    ]
                },
                "firstName": {
                    "type": "string"
                },
                "iat": {
                    "description": "the ` + "`" + `iat` + "`" + ` (Issued At) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.6",
                    "allOf": [
                        {
                            "$ref": "#/definitions/jwt.NumericDate"
                        }
                    ]
                },
                "id": {
                    "type": "string"
                },
                "iss": {
                    "description": "the ` + "`" + `iss` + "`" + ` (Issuer) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1",
                    "type": "string"
                },
                "jti": {
                    "description": "the ` + "`" + `jti` + "`" + ` (JWT ID) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.7",
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "nbf": {
                    "description": "the ` + "`" + `nbf` + "`" + ` (Not Before) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.5",
                    "allOf": [
                        {
                            "$ref": "#/definitions/jwt.NumericDate"
                        }
                    ]
                },
                "password": {
                    "type": "string"
                },
                "sub": {
                    "description": "the ` + "`" + `sub` + "`" + ` (Subject) claim. See https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.2",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "When the record was updated",
                    "type": "string"
                },
                "updatedBy": {
                    "description": "Who updated the record",
                    "type": "string"
                }
            }
        },
        "controllers.AuthGetResponse": {
            "type": "object",
            "required": [
                "claims"
            ],
            "properties": {
                "claims": {
                    "$ref": "#/definitions/auth.Claims"
                }
            }
        },
        "controllers.AuthPostRequest": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.AuthPostResponse": {
            "type": "object",
            "required": [
                "tokens",
                "user"
            ],
            "properties": {
                "tokens": {
                    "$ref": "#/definitions/controllers.AuthTokens"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "controllers.AuthPostSignInRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.AuthTokens": {
            "type": "object",
            "required": [
                "authorization"
            ],
            "properties": {
                "authorization": {
                    "type": "string"
                }
            }
        },
        "controllers.BooksDeleteResponse": {
            "type": "object",
            "required": [
                "bookId"
            ],
            "properties": {
                "bookId": {
                    "type": "string"
                }
            }
        },
        "controllers.BooksGetResponse": {
            "type": "object",
            "required": [
                "book"
            ],
            "properties": {
                "book": {
                    "$ref": "#/definitions/models.Book"
                }
            }
        },
        "controllers.BooksListResponse": {
            "type": "object",
            "required": [
                "books"
            ],
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Book"
                    }
                }
            }
        },
        "controllers.BooksPostRequest": {
            "type": "object",
            "required": [
                "author",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "finishedAt": {
                    "type": "string"
                },
                "purchasedAt": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "controllers.UsersGetResponse": {
            "type": "object",
            "required": [
                "user"
            ],
            "properties": {
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "controllers.UsersPostRequest": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "httputil.HTTPError": {
            "type": "object",
            "required": [
                "error"
            ],
            "properties": {
                "error": {
                    "$ref": "#/definitions/httputil.HTTPErrorError"
                }
            }
        },
        "httputil.HTTPErrorError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "jwt.NumericDate": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "models.Book": {
            "description": "A book",
            "type": "object",
            "required": [
                "author",
                "createdAt",
                "createdBy",
                "id",
                "title",
                "updatedAt",
                "updatedBy",
                "userId"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "createdAt": {
                    "description": "When the record was created",
                    "type": "string"
                },
                "createdBy": {
                    "description": "Who created the record",
                    "type": "string"
                },
                "deletedAt": {
                    "description": "When the record was deleted",
                    "type": "string"
                },
                "deletedBy": {
                    "description": "Who deleted the record",
                    "type": "string"
                },
                "finishedAt": {
                    "description": "The date the book was finished",
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "purchasedAt": {
                    "description": "The date the book was purchased",
                    "type": "string"
                },
                "rating": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 0
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "description": "When the record was updated",
                    "type": "string"
                },
                "updatedBy": {
                    "description": "Who updated the record",
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "description": "A user",
            "type": "object",
            "required": [
                "createdAt",
                "createdBy",
                "email",
                "emailVerified",
                "firstName",
                "id",
                "lastName",
                "updatedAt",
                "updatedBy"
            ],
            "properties": {
                "createdAt": {
                    "description": "When the record was created",
                    "type": "string"
                },
                "createdBy": {
                    "description": "Who created the record",
                    "type": "string"
                },
                "deletedAt": {
                    "description": "When the record was deleted",
                    "type": "string"
                },
                "deletedBy": {
                    "description": "Who deleted the record",
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "emailVerified": {
                    "type": "boolean"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updatedAt": {
                    "description": "When the record was updated",
                    "type": "string"
                },
                "updatedBy": {
                    "description": "Who updated the record",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Bearer token for access control",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "PMS API",
	Description:      "This API contains CRUD operations for PMS.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
