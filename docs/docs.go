// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://example.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://example.com/support",
            "email": "support@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/notifications": {
            "get": {
                "description": "Retrieve all notifications, optionally filtered by status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Get all notifications",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Notification status (e.g., unread)",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of notifications retrieved successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Notification"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Failed to fetch notifications",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new notification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Create a new notification",
                "parameters": [
                    {
                        "description": "Notification payload",
                        "name": "notification",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateNotification"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Notification created successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Notification"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid payload",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create notification",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/notifications/mark-all-read": {
            "put": {
                "description": "Mark all notifications as read",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Mark all notifications as read",
                "responses": {
                    "200": {
                        "description": "All notifications marked as read successfully",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to mark notifications as read",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/notifications/{id}": {
            "get": {
                "description": "Retrieve a notification by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Get a notification by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Notification ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Notification retrieved successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Notification"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid ID format",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch notification",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the status or details of a notification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Update a notification by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Notification ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Notification updated successfully",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid ID format",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to update notification",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a notification using its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Notification"
                ],
                "summary": "Delete a notification by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Notification ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Notification deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/model.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid ID format",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to delete notification",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/revenue/month": {
            "get": {
                "description": "Get total revenue grouped by month",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Revenue"
                ],
                "summary": "Fetch monthly revenue",
                "responses": {
                    "200": {
                        "description": "Fetch monthly revenue successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.MonthlyRevenue"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Failed to fetch monthly revenue",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/revenue/products": {
            "get": {
                "description": "Get revenue details for all products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Revenue"
                ],
                "summary": "Fetch product revenues",
                "responses": {
                    "200": {
                        "description": "Fetch product revenues successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.ProductRevenue"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Failed to fetch product revenues",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/revenue/status": {
            "get": {
                "description": "Get total revenue grouped by order status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Revenue"
                ],
                "summary": "Fetch total revenue by status",
                "responses": {
                    "200": {
                        "description": "Fetch total revenue by status successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.SuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.RevenueByStatus"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Failed to fetch total revenue by status",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateNotification": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "You have a new message"
                },
                "status": {
                    "type": "string",
                    "example": "new"
                },
                "title": {
                    "type": "string",
                    "example": "New Message"
                }
            }
        },
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.MonthlyRevenue": {
            "type": "object",
            "required": [
                "month",
                "revenue"
            ],
            "properties": {
                "month": {
                    "type": "string",
                    "example": "January"
                },
                "revenue": {
                    "type": "number",
                    "example": 100.5
                }
            }
        },
        "model.Notification": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "message": {
                    "type": "string",
                    "example": "You have a new message"
                },
                "status": {
                    "type": "string",
                    "example": "new"
                },
                "title": {
                    "type": "string",
                    "example": "New Message"
                }
            }
        },
        "model.ProductRevenue": {
            "type": "object",
            "required": [
                "product_name",
                "profit",
                "profit_margin",
                "revenue_date",
                "sell_price",
                "total_revenue"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "product_name": {
                    "type": "string",
                    "example": "Chicken Parmesan"
                },
                "profit": {
                    "type": "number",
                    "example": 7985
                },
                "profit_margin": {
                    "type": "number",
                    "example": 15
                },
                "revenue_date": {
                    "type": "string",
                    "example": "2024-03-28"
                },
                "sell_price": {
                    "type": "number",
                    "example": 55
                },
                "total_revenue": {
                    "type": "number",
                    "example": 8000
                }
            }
        },
        "model.RevenueByStatus": {
            "type": "object",
            "required": [
                "revenue",
                "status"
            ],
            "properties": {
                "revenue": {
                    "type": "number",
                    "example": 100.5
                },
                "status": {
                    "type": "string",
                    "example": "confirmed"
                }
            }
        },
        "model.SuccessResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authentication": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "UserID": {
            "type": "apiKey",
            "name": "User-ID",
            "in": "header"
        },
        "UserRole": {
            "type": "apiKey",
            "name": "User-Role",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "App pos Team 2",
	Description:      "API for Point Of Sale",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
