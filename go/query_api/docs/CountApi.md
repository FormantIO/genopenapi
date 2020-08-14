# \CountApi

All URIs are relative to *https://api.formant.io/v1/queries*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountControllerActiveDevices**](CountApi.md#CountControllerActiveDevices) | **Post** /counts/active-devices | Active devices
[**CountControllerHistory**](CountApi.md#CountControllerHistory) | **Post** /counts/history | History



## CountControllerActiveDevices

> UuidListResponse CountControllerActiveDevices(ctx, activeDevicesQuery)

Active devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activeDevicesQuery** | [**ActiveDevicesQuery**](ActiveDevicesQuery.md)| ActiveDevicesQuery | 

### Return type

[**UuidListResponse**](UuidListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountControllerHistory

> CountHistory CountControllerHistory(ctx, countHistoryQuery)

History

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countHistoryQuery** | [**CountHistoryQuery**](CountHistoryQuery.md)| CountHistoryQuery | 

### Return type

[**CountHistory**](CountHistory.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

