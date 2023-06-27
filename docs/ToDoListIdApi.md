# {{classname}}

All URIs are relative to *https://{environment}.programmerzamannow.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TodolistTodolistIdDelete**](ToDoListIdApi.md#TodolistTodolistIdDelete) | **Delete** /todolist/{todolistId} | Delete Existing ToDoList
[**TodolistTodolistIdPut**](ToDoListIdApi.md#TodolistTodolistIdPut) | **Put** /todolist/{todolistId} | Update Existing Todolist

# **TodolistTodolistIdDelete**
> InlineResponse2001 TodolistTodolistIdDelete(ctx, todolistId)
Delete Existing ToDoList

Delete Existing Todolist In Database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **todolistId** | **string**| ToDoList Id For Updated | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ToDoListAuth](../README.md#ToDoListAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TodolistTodolistIdPut**
> []Todolist TodolistTodolistIdPut(ctx, body, todolistId)
Update Existing Todolist

Update Existing Todolist In Database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TodolistTodolistIdBody**](TodolistTodolistIdBody.md)|  | 
  **todolistId** | **string**| ToDoList Id For Updated | 

### Return type

[**[]Todolist**](array.md)

### Authorization

[ToDoListAuth](../README.md#ToDoListAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

