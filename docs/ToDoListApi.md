# {{classname}}

All URIs are relative to *https://{environment}.programmerzamannow.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TodolistGet**](ToDoListApi.md#TodolistGet) | **Get** /todolist | Get All ToDoList
[**TodolistPost**](ToDoListApi.md#TodolistPost) | **Post** /todolist | Create New ToDoList

# **TodolistGet**
> []InlineResponse200 TodolistGet(ctx, optional)
Get All ToDoList

Get All Active ToDoList

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ToDoListApiTodolistGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ToDoListApiTodolistGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDone** | **optional.Bool**| Is Include Done ToDoList | [default to true]
 **name** | **optional.String**| Filter ToDoList By Name | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[ToDoListAuth](../README.md#ToDoListAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TodolistPost**
> []InlineResponse200 TodolistPost(ctx, body)
Create New ToDoList

Create New ToDoList To Database

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrUpdateTodolist**](CreateOrUpdateTodolist.md)|  | 

### Return type

[**[]InlineResponse200**](inline_response_200.md)

### Authorization

[ToDoListAuth](../README.md#ToDoListAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

