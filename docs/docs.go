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
        "/api/v1/auth": {
            "get": {
                "description": "获取token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"token\":\"xx\"},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "description": "获取标签列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签接口"
                ],
                "summary": "列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标签名",
                        "name": "name",
                        "in": "path"
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "state",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "新增标签",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "新增",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标签名",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "state",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "创建人",
                        "name": "created_by",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/:id": {
            "put": {
                "description": "编辑标签",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "编辑",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标签名",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "状态",
                        "name": "state",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "修改人",
                        "name": "modified_by",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
