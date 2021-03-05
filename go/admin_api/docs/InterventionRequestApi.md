# \InterventionRequestApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InterventionRequestControllerGetOne**](InterventionRequestApi.md#InterventionRequestControllerGetOne) | **Get** /intervention-requests/{id} | Get one
[**InterventionRequestControllerPost**](InterventionRequestApi.md#InterventionRequestControllerPost) | **Post** /intervention-requests | Post



## InterventionRequestControllerGetOne

> InterventionRequest InterventionRequestControllerGetOne(ctx, id)

Get one

Get an intervention request Authorized clients: viewer, device Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**InterventionRequest**](InterventionRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InterventionRequestControllerPost

> InterventionRequest InterventionRequestControllerPost(ctx, interventionRequest)

Post

Create an intervention request Authorized clients: device Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interventionRequest** | [**InterventionRequest**](InterventionRequest.md)| InterventionRequest | 

### Return type

[**InterventionRequest**](InterventionRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

