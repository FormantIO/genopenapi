# \CommandResponseApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandResponseControllerPost**](CommandResponseApi.md#CommandResponseControllerPost) | **Post** /command-responses | Post



## CommandResponseControllerPost

> CommandResponse CommandResponseControllerPost(ctx, commandResponse)

Post

Respond to a command

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandResponse** | [**CommandResponse**](CommandResponse.md)| CommandResponse | 

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

