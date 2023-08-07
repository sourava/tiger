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
        "/login": {
            "post": {
                "description": "provides token if email password combination is correct",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "login api",
                "parameters": [
                    {
                        "description": "login request body params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_business_auth_request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_business_auth_response.LoginHandlerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_external_utils.HandlerErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_external_utils.HandlerErrorResponse"
                        }
                    }
                }
            }
        },
        "/tigers": {
            "get": {
                "description": "returns all tigers sorted by last time the tiger was seen.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "list all tigers api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_business_tiger_response.ListAllTigersHandlerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_external_utils.HandlerErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_external_utils.HandlerErrorResponse"
                        }
                    }
                }
            }
        },
        "/tigers/:tigerID/sightings": {
            "get": {
                "description": "returns all sightings for a tiger sorted by date.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "list all tiger sightings api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page Size",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Tiger ID",
                        "name": "tigerID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_business_tiger_response.ListAllTigersHandlerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_external_utils.HandlerErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_external_utils.HandlerErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "creates a user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "create user api",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token received in login api response",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "create user request body params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_business_user_request.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_business_user_response.CreateUserHandlerResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_external_utils.HandlerErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_sourava_tiger_external_utils.HandlerErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_sourava_tiger_business_auth_request.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@gmail.com"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                }
            }
        },
        "github_com_sourava_tiger_business_auth_response.LoginHandlerResponse": {
            "type": "object",
            "properties": {
                "payload": {
                    "$ref": "#/definitions/github_com_sourava_tiger_business_auth_response.LoginResponse"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "github_com_sourava_tiger_business_auth_response.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "jwt.token"
                }
            }
        },
        "github_com_sourava_tiger_business_tiger_response.ListAllTigersHandlerResponse": {
            "type": "object",
            "properties": {
                "payload": {
                    "$ref": "#/definitions/github_com_sourava_tiger_business_tiger_response.ListAllTigersResponse"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "github_com_sourava_tiger_business_tiger_response.ListAllTigersResponse": {
            "type": "object",
            "properties": {
                "tigers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_sourava_tiger_business_tiger_response.TigerResponse"
                    }
                }
            }
        },
        "github_com_sourava_tiger_business_tiger_response.TigerResponse": {
            "type": "object",
            "properties": {
                "Last_seen_longitude": {
                    "type": "number",
                    "example": 10.2
                },
                "date_of_birth": {
                    "type": "string",
                    "example": "2020-01-13"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "last_seen_latitude": {
                    "type": "number",
                    "example": 1.1
                },
                "last_seen_timestamp": {
                    "type": "integer",
                    "example": 1691354650
                },
                "name": {
                    "type": "string",
                    "example": "tiger name"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "github_com_sourava_tiger_business_user_request.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "github_com_sourava_tiger_business_user_response.CreateUserHandlerResponse": {
            "type": "object",
            "properties": {
                "payload": {
                    "$ref": "#/definitions/github_com_sourava_tiger_business_user_response.CreateUserResponse"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "github_com_sourava_tiger_business_user_response.CreateUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@gmail.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "github_com_sourava_tiger_external_utils.HandlerErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Something went wrong"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Tigerhall Kittens API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
