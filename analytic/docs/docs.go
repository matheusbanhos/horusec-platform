// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Horusec",
            "url": "https://github.com/ZupIT/horusec-platform",
            "email": "horusec@zup.com.br"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/analytic/dashboard/{workspaceID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all charts of dashboard screen",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dashboard"
                ],
                "operationId": "GetAllChartsByWorkspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "workspaceID of the workspace",
                        "name": "workspaceID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "initialDate query string",
                        "name": "initialDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "finalDate query string",
                        "name": "finalDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "$ref": "#/definitions/dashboard.Response"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "BAD REQUEST",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "INTERNAL SERVER ERROR",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/analytic/dashboard/{workspaceID}/{repositoryID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all charts of dashboard screen",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dashboard"
                ],
                "operationId": "GetAllChartsByRepository",
                "parameters": [
                    {
                        "type": "string",
                        "description": "workspaceID of the workspace",
                        "name": "workspaceID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "repositoryID of the repository",
                        "name": "repositoryID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "initialDate query string",
                        "name": "initialDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "finalDate query string",
                        "name": "finalDate",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "$ref": "#/definitions/dashboard.Response"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "BAD REQUEST",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "INTERNAL SERVER ERROR",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/analytic/health": {
            "get": {
                "description": "Check if Health of service it's OK!",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "operationId": "health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "INTERNAL SERVER ERROR",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "content": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dashboard.Response": {
            "type": "object",
            "properties": {
                "totalAuthors": {
                    "type": "integer"
                },
                "totalRepositories": {
                    "type": "integer"
                },
                "vulnerabilitiesByAuthor": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dashboard.ResponseByAuthor"
                    }
                },
                "vulnerabilitiesByLanguage": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dashboard.ResponseByLanguage"
                    }
                },
                "vulnerabilitiesByRepository": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dashboard.ResponseByRepository"
                    }
                },
                "vulnerabilityBySeverity": {
                    "$ref": "#/definitions/dashboard.BySeverities"
                },
                "vulnerabilityByTime": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dashboard.ResponseByTime"
                    }
                }
            }
        },
        "dashboard.ResponseByAuthor": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "critical": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "high": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "info": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "low": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "medium": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "unknown": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                }
            }
        },
        "dashboard.ResponseByLanguage": {
            "type": "object",
            "properties": {
                "critical": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "high": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "info": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "language": {
                    "type": "string"
                },
                "low": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "medium": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "unknown": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                }
            }
        },
        "dashboard.ResponseByRepository": {
            "type": "object",
            "properties": {
                "critical": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "high": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "info": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "low": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "medium": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "repositoryName": {
                    "type": "string"
                },
                "unknown": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                }
            }
        },
        "dashboard.ResponseByTime": {
            "type": "object",
            "properties": {
                "critical": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "high": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "info": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "low": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "medium": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "time": {
                    "type": "string"
                },
                "unknown": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                }
            }
        },
        "dashboard.BySeverities": {
            "type": "object",
            "properties": {
                "critical": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "high": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "info": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "low": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "medium": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                },
                "unknown": {
                    "$ref": "#/definitions/dashboard.ResponseSeverityContAndTypes"
                }
            }
        },
        "dashboard.ResponseSeverityContAndTypes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "types": {
                    "$ref": "#/definitions/dashboard.ResponseVulnTypes"
                }
            }
        },
        "dashboard.ResponseVulnTypes": {
            "type": "object",
            "properties": {
                "corrected": {
                    "type": "integer"
                },
                "falsePositive": {
                    "type": "integer"
                },
                "riskAccepted": {
                    "type": "integer"
                },
                "vulnerability": {
                    "type": "integer"
                }
            }
        },
        "entities.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "content": {
                    "type": "object"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "X-Horusec-Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Horusec-Analytic",
	Description: "Service responsible for managing vulnerabilities.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
