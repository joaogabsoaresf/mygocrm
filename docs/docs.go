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
        "/lead": {
            "get": {
                "description": "Show a specifc lead in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Leads"
                ],
                "summary": "Show lead",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lead id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ShowLeadResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a leads in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Leads"
                ],
                "summary": "Update lead",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateLeadRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Lead id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateLeadResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new lead in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Leads"
                ],
                "summary": "Create lead",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateLeadRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateLeadResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a lead in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Leads"
                ],
                "summary": "Delete lead",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lead id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.DeleteLeadResponse"
                        }
                    }
                }
            }
        },
        "/leads": {
            "get": {
                "description": "Show all leads in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Leads"
                ],
                "summary": "Show leads",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ShowLeadResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateLeadRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "client": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "handler.CreateLeadResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.LeadRespose"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.DeleteLeadResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.LeadRespose"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ShowLeadResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.LeadRespose"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateLeadRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "client": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateLeadResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.LeadRespose"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.LeadRespose": {
            "type": "object",
            "properties": {
                "client": {
                    "type": "boolean"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
