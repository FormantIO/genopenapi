# \TriggeredEventApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TriggeredEventControllerGetOne**](TriggeredEventApi.md#TriggeredEventControllerGetOne) | **Get** /triggered-events/{id} | Get one
[**TriggeredEventControllerPost**](TriggeredEventApi.md#TriggeredEventControllerPost) | **Post** /triggered-events | Post



## TriggeredEventControllerGetOne

> TriggeredEvent TriggeredEventControllerGetOne(ctx, id)

Get one

Get an triggered event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**TriggeredEvent**](TriggeredEvent.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggeredEventControllerPost

> TriggeredEvent TriggeredEventControllerPost(ctx, triggeredEvent)

Post

Create an triggeredEvent

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**triggeredEvent** | [**TriggeredEvent**](TriggeredEvent.md)| TriggeredEvent | 

### Return type

[**TriggeredEvent**](TriggeredEvent.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

