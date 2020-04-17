# \BoardApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BoardControllerDelete**](BoardApi.md#BoardControllerDelete) | **Delete** /boards/{id} | Delete
[**BoardControllerGetAll**](BoardApi.md#BoardControllerGetAll) | **Get** /boards | Get all
[**BoardControllerGetOne**](BoardApi.md#BoardControllerGetOne) | **Get** /boards/{id} | Get one
[**BoardControllerPatch**](BoardApi.md#BoardControllerPatch) | **Patch** /boards/{id} | Patch
[**BoardControllerPost**](BoardApi.md#BoardControllerPost) | **Post** /boards | Post



## BoardControllerDelete

> BoardControllerDelete(ctx, id)

Delete

Delete a board

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardControllerGetAll

> BoardListResponse BoardControllerGetAll(ctx, )

Get all

List all boards

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**BoardListResponse**](BoardListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardControllerGetOne

> Board BoardControllerGetOne(ctx, id)

Get one

Get a board

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Board**](Board.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardControllerPatch

> Board BoardControllerPatch(ctx, id, partialBoard)

Patch

Update a board

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialBoard** | [**PartialBoard**](PartialBoard.md)| PartialBoard | 

### Return type

[**Board**](Board.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BoardControllerPost

> Board BoardControllerPost(ctx, board)

Post

Create board

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**board** | [**Board**](Board.md)| Board | 

### Return type

[**Board**](Board.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

