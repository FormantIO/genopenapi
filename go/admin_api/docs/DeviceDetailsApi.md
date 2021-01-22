# \DeviceDetailsApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceDetailsControllerGetAll**](DeviceDetailsApi.md#DeviceDetailsControllerGetAll) | **Post** /device-details/query | Get all
[**DeviceDetailsControllerGetOne**](DeviceDetailsApi.md#DeviceDetailsControllerGetOne) | **Get** /device-details/{id} | Get one



## DeviceDetailsControllerGetAll

> DeviceDetailsListResponse DeviceDetailsControllerGetAll(ctx, optional)

Get all

Authorized clients: viewer Authorized plans: freemium, commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceDetailsControllerGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceDetailsControllerGetAllOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceQuery** | [**optional.Interface of DeviceQuery**](DeviceQuery.md)| DeviceQuery | 

### Return type

[**DeviceDetailsListResponse**](DeviceDetailsListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceDetailsControllerGetOne

> DeviceDetails DeviceDetailsControllerGetOne(ctx, id)

Get one

Authorized clients: viewer Authorized plans: freemium, commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**DeviceDetails**](DeviceDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

