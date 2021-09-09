# \ShareApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShareControllerAuthenticate**](ShareApi.md#ShareControllerAuthenticate) | **Post** /shares/{code}/authenticate | Authenticate
[**ShareControllerDelete**](ShareApi.md#ShareControllerDelete) | **Delete** /shares/{code} | Delete
[**ShareControllerGetAll**](ShareApi.md#ShareControllerGetAll) | **Get** /shares | Get all
[**ShareControllerGetOne**](ShareApi.md#ShareControllerGetOne) | **Get** /shares/{code} | Get one
[**ShareControllerPost**](ShareApi.md#ShareControllerPost) | **Post** /shares | Post



## ShareControllerAuthenticate

> TokenResult ShareControllerAuthenticate(ctx, code)

Authenticate

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**|  | 

### Return type

[**TokenResult**](TokenResult.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareControllerDelete

> ShareControllerDelete(ctx, code)

Delete

Revoke a share Authorized clients: operator Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareControllerGetAll

> ShareListResponse ShareControllerGetAll(ctx, )

Get all

List shares Authorized clients: operator Authorized plans: starter, premium

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ShareListResponse**](ShareListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareControllerGetOne

> Share ShareControllerGetOne(ctx, code)

Get one

Get a share

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**|  | 

### Return type

[**Share**](Share.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareControllerPost

> Share ShareControllerPost(ctx, share)

Post

Create a share token Authorized clients: operator Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**share** | [**Share**](Share.md)| Share | 

### Return type

[**Share**](Share.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

