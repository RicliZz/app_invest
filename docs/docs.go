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
        "/admin/as-user/{id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Только админ",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Зайти под пользователем",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя, под которым надо войти",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT токен",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Неверные параметры",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/get-all": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "ФИО/Название стартапа",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Найти всех пользователей по параметрам",
                "parameters": [
                    {
                        "description": "Параметры поиска",
                        "name": "searchParams",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/AdminService.SearchParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Найденные пользователи",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/userModel.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Неверные параметры",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/search-one/{id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Найти пользователя по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Найденный юзер",
                        "schema": {
                            "$ref": "#/definitions/userModel.User"
                        }
                    },
                    "400": {
                        "description": "Пользователь по ID не найден",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "Вход в систему",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "sign-in",
                "parameters": [
                    {
                        "description": "Данные для входа",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authModel.RequestSignInPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT токен",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "Регистрация в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "sign-up",
                "parameters": [
                    {
                        "description": "Данные для регистрации",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authModel.RequestSignUpPayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/info/about-us": {
            "get": {
                "description": "Информация о нас",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Info"
                ],
                "summary": "about-us",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/info/contacts": {
            "get": {
                "description": "Контакты для связи",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Info"
                ],
                "summary": "contacts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/profile/delete": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Удалить профиль",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "deleteProfile",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/profile/statistic": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Посмотреть свой профиль",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "getStatistic",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Need int ID",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/profile/update": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Обновить профиль",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "patchProfile",
                "parameters": [
                    {
                        "description": "Данные для обновления профиля",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.PatchProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/profile/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Посмотреть чужой профиль",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "getProfile",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Need int ID",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startup/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Создание стартапа",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StartUp"
                ],
                "summary": "create startup",
                "parameters": [
                    {
                        "description": "Данные для создания стартапа",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.StartupRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startup/most-popular": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Топ-5 стартапов",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StartUp"
                ],
                "summary": "Получить ТОП-5 стартапов",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startup/search-all": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Просмотреть все стартапы",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StartUp"
                ],
                "summary": "getAllStartups",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startup/{id}/": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Просмотреть стартап",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StartUp"
                ],
                "summary": "getStartup",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID стартапа",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startup/{id}/delete": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Удаление стартапа",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StartUp"
                ],
                "summary": "deleteStartup",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID стартапа",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/startup/{id}/update": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Обновление стартапа",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "StartUp"
                ],
                "summary": "update startup",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID стартапа",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные для обновления стартапа",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.StartupRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "AdminService.SearchParams": {
            "type": "object",
            "properties": {
                "HaveStartups": {
                    "type": "boolean",
                    "example": false
                },
                "NotConfirmed": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "authModel.RequestSignInPayload": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "admin@mail.ru"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "12345678"
                }
            }
        },
        "authModel.RequestSignUpPayload": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "middleName",
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
                "middleName": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "enums.StartUpStage": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "IDEA",
                "PROTOTYPE",
                "PRODUCT"
            ]
        },
        "enums.StartUpStatus": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "CLOSE",
                "ACTIVE"
            ]
        },
        "enums.StartUpTopic": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "SPORT",
                "TEACHING",
                "CAR"
            ]
        },
        "requests.PatchProfileRequest": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/requests.UpdateUserRequest"
                },
                "user_details": {
                    "$ref": "#/definitions/requests.UpdateUserDetailsRequest"
                }
            }
        },
        "requests.StartupRequest": {
            "type": "object",
            "properties": {
                "fundingGoal": {
                    "description": "required",
                    "type": "number"
                },
                "historyOfCreation": {
                    "type": "string"
                },
                "idea": {
                    "type": "string"
                },
                "offeredPercent": {
                    "description": "required",
                    "type": "number"
                },
                "stage": {
                    "$ref": "#/definitions/enums.StartUpStage"
                },
                "status": {
                    "description": "required",
                    "allOf": [
                        {
                            "$ref": "#/definitions/enums.StartUpStatus"
                        }
                    ]
                },
                "strategy": {
                    "type": "string"
                },
                "title": {
                    "description": "required",
                    "type": "string"
                },
                "topic": {
                    "description": "required",
                    "allOf": [
                        {
                            "$ref": "#/definitions/enums.StartUpTopic"
                        }
                    ]
                }
            }
        },
        "requests.UpdateUserDetailsRequest": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                }
            }
        },
        "requests.UpdateUserRequest": {
            "type": "object",
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
                "middleName": {
                    "type": "string"
                }
            }
        },
        "responses.UserDetailsResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "socials": {
                    "$ref": "#/definitions/userModel.Socials"
                }
            }
        },
        "responses.UserResponse": {
            "type": "object",
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
                "middleName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "userDetails": {
                    "$ref": "#/definitions/responses.UserDetailsResponse"
                }
            }
        },
        "startUpModel.FounderSocials": {
            "type": "object",
            "properties": {
                "instagramLink": {
                    "type": "string"
                },
                "telegramLink": {
                    "type": "string"
                },
                "vkLink": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                },
                "whatsUpLink": {
                    "type": "string"
                }
            }
        },
        "startUpModel.StartUp": {
            "type": "object",
            "required": [
                "fundingGoal",
                "offeredPercent",
                "status",
                "title",
                "topic"
            ],
            "properties": {
                "createdAt": {
                    "description": "Информация о дате создания и апдейта",
                    "type": "string"
                },
                "founderEmail": {
                    "type": "string"
                },
                "founderFullName": {
                    "description": "Информация о создателе",
                    "type": "string"
                },
                "founderSocials": {
                    "$ref": "#/definitions/startUpModel.FounderSocials"
                },
                "fundingGoal": {
                    "description": "цель по финансированию, сколько нужно собрать",
                    "type": "number"
                },
                "historyOfCreation": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "idea": {
                    "type": "string"
                },
                "offeredPercent": {
                    "description": "предлагаемый процент инвестору",
                    "type": "number"
                },
                "stage": {
                    "$ref": "#/definitions/enums.StartUpStage"
                },
                "status": {
                    "$ref": "#/definitions/enums.StartUpStatus"
                },
                "strategy": {
                    "type": "string"
                },
                "title": {
                    "description": "Информация о стартапе",
                    "type": "string"
                },
                "topic": {
                    "$ref": "#/definitions/enums.StartUpTopic"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "userModel.Socials": {
            "type": "object",
            "properties": {
                "instagramLink": {
                    "type": "string"
                },
                "telegramLink": {
                    "type": "string"
                },
                "vkLink": {
                    "type": "string"
                },
                "website": {
                    "type": "string"
                },
                "whatsUpLink": {
                    "type": "string"
                }
            }
        },
        "userModel.User": {
            "type": "object",
            "properties": {
                "aboutMe": {
                    "type": "string"
                },
                "confirmed": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "emailConfirm": {
                    "type": "boolean"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "middleName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "startUps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/startUpModel.StartUp"
                    }
                },
                "userDetails": {
                    "$ref": "#/definitions/userModel.UserDetails"
                }
            }
        },
        "userModel.UserDetails": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "socials": {
                    "$ref": "#/definitions/userModel.Socials"
                },
                "userId": {
                    "type": "integer"
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
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "HaveIdea",
	Description:      "This is a sample server for demonstrating JWT with Swagger.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
