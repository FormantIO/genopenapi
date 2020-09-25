# \ShareApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShareControllerAuthenticate**](ShareApi.md#ShareControllerAuthenticate) | **Post** /shares/{code}/authenticate | Authenticate
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

Create a share token Authorized clients: operator

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

