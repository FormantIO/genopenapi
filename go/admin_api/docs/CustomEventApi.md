# \CustomEventApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomEventControllerPost**](CustomEventApi.md#CustomEventControllerPost) | **Post** /custom-events | Post



## CustomEventControllerPost

> CustomEvent CustomEventControllerPost(ctx, customEvent)

Post

Create a custom event Authorized clients: device Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customEvent** | [**CustomEvent**](CustomEvent.md)| CustomEvent | 

### Return type

[**CustomEvent**](CustomEvent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

