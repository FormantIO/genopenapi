# \CaptureSessionApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CaptureSessionControllerAuthenticate**](CaptureSessionApi.md#CaptureSessionControllerAuthenticate) | **Post** /capture-sessions/{code}/authenticate | Authenticate
[**CaptureSessionControllerGetOne**](CaptureSessionApi.md#CaptureSessionControllerGetOne) | **Get** /capture-sessions/{code} | Get one
[**CaptureSessionControllerPost**](CaptureSessionApi.md#CaptureSessionControllerPost) | **Post** /capture-sessions | Post



## CaptureSessionControllerAuthenticate

> TokenResult CaptureSessionControllerAuthenticate(ctx, code)

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


## CaptureSessionControllerGetOne

> CaptureSession CaptureSessionControllerGetOne(ctx, code)

Get one

Get a capture session

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**|  | 

### Return type

[**CaptureSession**](CaptureSession.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CaptureSessionControllerPost

> CaptureSession CaptureSessionControllerPost(ctx, captureSession)

Post

Create a capture session Authorized clients: operator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**captureSession** | [**CaptureSession**](CaptureSession.md)| CaptureSession | 

### Return type

[**CaptureSession**](CaptureSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

