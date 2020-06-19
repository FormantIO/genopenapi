# \EventTriggerApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventTriggerControllerGetAll**](EventTriggerApi.md#EventTriggerControllerGetAll) | **Get** /event-triggers | Get all
[**EventTriggerControllerGetOne**](EventTriggerApi.md#EventTriggerControllerGetOne) | **Get** /event-triggers/{id} | Get one
[**EventTriggerControllerPatch**](EventTriggerApi.md#EventTriggerControllerPatch) | **Patch** /event-triggers/{id} | Patch
[**EventTriggerControllerPost**](EventTriggerApi.md#EventTriggerControllerPost) | **Post** /event-triggers | Post



## EventTriggerControllerGetAll

> EventTriggerListResponse EventTriggerControllerGetAll(ctx, )

Get all

List all event triggers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EventTriggerListResponse**](EventTriggerListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventTriggerControllerGetOne

> EventTrigger EventTriggerControllerGetOne(ctx, id)

Get one

Get an event trigger

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**EventTrigger**](EventTrigger.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventTriggerControllerPatch

> EventTrigger EventTriggerControllerPatch(ctx, id, partialEventTrigger)

Patch

Update event trigger

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialEventTrigger** | [**PartialEventTrigger**](PartialEventTrigger.md)| PartialEventTrigger | 

### Return type

[**EventTrigger**](EventTrigger.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventTriggerControllerPost

> EventTrigger EventTriggerControllerPost(ctx, eventTrigger)

Post

Create event trigger

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTrigger** | [**EventTrigger**](EventTrigger.md)| EventTrigger | 

### Return type

[**EventTrigger**](EventTrigger.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

