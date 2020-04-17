# \AlertApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertControllerGetOne**](AlertApi.md#AlertControllerGetOne) | **Get** /alerts/{id} | Get one
[**AlertControllerPost**](AlertApi.md#AlertControllerPost) | **Post** /alerts | Post



## AlertControllerGetOne

> Alert AlertControllerGetOne(ctx, id)

Get one

Get an alert

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Alert**](Alert.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertControllerPost

> Alert AlertControllerPost(ctx, alert)

Post

Create an alert

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alert** | [**Alert**](Alert.md)| Alert | 

### Return type

[**Alert**](Alert.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

