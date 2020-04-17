# \DeviceLayoutApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceLayoutControllerDelete**](DeviceLayoutApi.md#DeviceLayoutControllerDelete) | **Delete** /device-layouts/{id} | Delete
[**DeviceLayoutControllerGetAll**](DeviceLayoutApi.md#DeviceLayoutControllerGetAll) | **Get** /device-layouts | Get all
[**DeviceLayoutControllerGetOne**](DeviceLayoutApi.md#DeviceLayoutControllerGetOne) | **Get** /device-layouts/{id} | Get one
[**DeviceLayoutControllerPatch**](DeviceLayoutApi.md#DeviceLayoutControllerPatch) | **Patch** /device-layouts/{id} | Patch
[**DeviceLayoutControllerPost**](DeviceLayoutApi.md#DeviceLayoutControllerPost) | **Post** /device-layouts | Post



## DeviceLayoutControllerDelete

> DeviceLayoutControllerDelete(ctx, id)

Delete

Delete a device layout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceLayoutControllerGetAll

> DeviceListResponse DeviceLayoutControllerGetAll(ctx, )

Get all

List all device layouts

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DeviceListResponse**](DeviceListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceLayoutControllerGetOne

> DeviceLayout DeviceLayoutControllerGetOne(ctx, id)

Get one

Get a device layout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**DeviceLayout**](DeviceLayout.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceLayoutControllerPatch

> DeviceLayout DeviceLayoutControllerPatch(ctx, id, partialDeviceLayout)

Patch

Update a device layout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialDeviceLayout** | [**PartialDeviceLayout**](PartialDeviceLayout.md)| PartialDeviceLayout | 

### Return type

[**DeviceLayout**](DeviceLayout.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceLayoutControllerPost

> DeviceLayout DeviceLayoutControllerPost(ctx, deviceLayout)

Post

Create a device layout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceLayout** | [**DeviceLayout**](DeviceLayout.md)| DeviceLayout | 

### Return type

[**DeviceLayout**](DeviceLayout.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

