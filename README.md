# Go API client for swagger

OpenAPI for ToDoList Restful API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://www.programmerzamannow.com](https://www.programmerzamannow.com)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://{environment}.programmerzamannow.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ToDoListApi* | [**TodolistGet**](docs/ToDoListApi.md#todolistget) | **Get** /todolist | Get All ToDoList
*ToDoListApi* | [**TodolistPost**](docs/ToDoListApi.md#todolistpost) | **Post** /todolist | Create New ToDoList
*ToDoListIdApi* | [**TodolistTodolistIdDelete**](docs/ToDoListIdApi.md#todolisttodolistiddelete) | **Delete** /todolist/{todolistId} | Delete Existing ToDoList
*ToDoListIdApi* | [**TodolistTodolistIdPut**](docs/ToDoListIdApi.md#todolisttodolistidput) | **Put** /todolist/{todolistId} | Update Existing Todolist

## Documentation For Models

 - [CreateOrUpdateTodolist](docs/CreateOrUpdateTodolist.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [Todolist](docs/Todolist.md)
 - [TodolistTodolistIdBody](docs/TodolistTodolistIdBody.md)

## Documentation For Authorization

## ToDoListAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

mhdianrush@gmail.com
