// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/doctor/patient": {
            "post": {
                "description": "This endpoint allows a doctor to create a new patient record by providing details like first name, last name, age, doctor_id, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "Create a new patient record (Doctor Role)",
                "parameters": [
                    {
                        "description": "Patient data",
                        "name": "patient",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created patient",
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/doctor/patient/{id}": {
            "put": {
                "description": "This endpoint allows a doctor to update an existing patient record. You need to provide the patient's ID in the URL and the new details in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "Update an existing patient record by ID (Doctor Role)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated patient data",
                        "name": "patient",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated patient",
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Patient not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "This endpoint allows a doctor to delete a patient record by providing the patient's ID in the URL.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "Delete a patient record by ID (Doctor Role)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted patient",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Patient not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/doctor/patients": {
            "get": {
                "description": "This endpoint retrieves all patient records from the database and returns them as a JSON array. This operation does not require authentication.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "Retrieve a list of all patients",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved all patients",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Patient"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "This endpoint allows a user to log in by providing their email and password. If the credentials are correct, a JWT token is generated and sent as a cookie in the response.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Log in a user using email and password",
                "parameters": [
                    {
                        "description": "User login credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful, JWT token set in cookie",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Incorrect password",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/receptionist/patient": {
            "post": {
                "description": "This endpoint allows the creation of a new patient record by providing details like first name, last name, email, phone number, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "Create a new patient record",
                "parameters": [
                    {
                        "description": "Patient data",
                        "name": "patient",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created patient",
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/receptionist/patient/{id}": {
            "put": {
                "description": "This endpoint allows you to update an existing patient record. You need to provide the patient's ID in the URL and the new details in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "Update an existing patient record by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated patient data",
                        "name": "patient",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated patient",
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Patient not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "This endpoint allows you to delete a patient record by providing the patient's ID in the URL.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "Delete a patient record by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted patient",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Patient not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/receptionist/patients": {
            "get": {
                "description": "This endpoint retrieves all patient records from the database and returns them as a JSON array.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Patients"
                ],
                "summary": "Get all patient records",
                "responses": {
                    "200": {
                        "description": "Successfully retrieved all patients",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Patient"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "This endpoint allows you to register a new user by providing their username, email, password, and role. The password is hashed before storing it in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Register"
                ],
                "summary": "Register a new user by provided data (username, email, password, role)",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created user",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "409": {
                        "description": "Email already exists",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Patient": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "doctor_id": {
                    "type": "integer"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "doctor patient site",
	Description:      "Tis is a simple medical field.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}