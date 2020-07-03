# \CommandApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandControllerGetOne**](CommandApi.md#CommandControllerGetOne) | **Get** /commands/{id} | Get one
[**CommandControllerPatch**](CommandApi.md#CommandControllerPatch) | **Patch** /commands/{id} | Patch
[**CommandControllerPost**](CommandApi.md#CommandControllerPost) | **Post** /commands | Post
[**CommandControllerQuery**](CommandApi.md#CommandControllerQuery) | **Post** /commands/query | Query



## CommandControllerGetOne

> Command CommandControllerGetOne(ctx, id)

Get one

Get a command

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Command**](Command.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommandControllerPatch

> Command CommandControllerPatch(ctx, id, partialCommand)

Patch

Update command

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialCommand** | [**PartialCommand**](PartialCommand.md)| PartialCommand | 

### Return type

[**Command**](Command.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommandControllerPost

> Command CommandControllerPost(ctx, command)

Post

Create a command

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**command** | [**Command**](Command.md)| Command | 

### Return type

[**Command**](Command.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommandControllerQuery

> CommandListResponse CommandControllerQuery(ctx, commandQuery)

Query

Query undelivered commands by device ID

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandQuery** | [**CommandQuery**](CommandQuery.md)| CommandQuery | 

### Return type

[**CommandListResponse**](CommandListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

