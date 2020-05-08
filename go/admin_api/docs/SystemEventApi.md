# \SystemEventApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemEventControllerPost**](SystemEventApi.md#SystemEventControllerPost) | **Post** /system-events | Post



## SystemEventControllerPost

> SystemEvent SystemEventControllerPost(ctx, systemEvent)

Post

Create a system event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**systemEvent** | [**SystemEvent**](SystemEvent.md)| SystemEvent | 

### Return type

[**SystemEvent**](SystemEvent.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

