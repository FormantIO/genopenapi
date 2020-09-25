# \InterventionResponseApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InterventionResponseControllerPost**](InterventionResponseApi.md#InterventionResponseControllerPost) | **Post** /intervention-responses | Post



## InterventionResponseControllerPost

> InterventionResponse InterventionResponseControllerPost(ctx, interventionResponse)

Post

Create an intervention response Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interventionResponse** | [**InterventionResponse**](InterventionResponse.md)| InterventionResponse | 

### Return type

[**InterventionResponse**](InterventionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

