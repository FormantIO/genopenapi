# \TriggeredEventApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TriggeredEventControllerGetOne**](TriggeredEventApi.md#TriggeredEventControllerGetOne) | **Get** /triggered-events/{id} | Get one



## TriggeredEventControllerGetOne

> TriggeredEvent TriggeredEventControllerGetOne(ctx, id)

Get one

Get an triggered event Authorized clients: viewer Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**TriggeredEvent**](TriggeredEvent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

