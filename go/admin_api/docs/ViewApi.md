# \ViewApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewControllerDelete**](ViewApi.md#ViewControllerDelete) | **Delete** /views/{id} | Delete
[**ViewControllerGetAll**](ViewApi.md#ViewControllerGetAll) | **Get** /views | Get all
[**ViewControllerGetOne**](ViewApi.md#ViewControllerGetOne) | **Get** /views/{id} | Get one
[**ViewControllerPatch**](ViewApi.md#ViewControllerPatch) | **Patch** /views/{id} | Patch
[**ViewControllerPost**](ViewApi.md#ViewControllerPost) | **Post** /views | Post



## ViewControllerDelete

> ViewControllerDelete(ctx, id)

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


## ViewControllerGetAll

> DeviceListResponse ViewControllerGetAll(ctx, )

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


## ViewControllerGetOne

> View ViewControllerGetOne(ctx, id)

Get one

Get a device layout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**View**](View.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewControllerPatch

> View ViewControllerPatch(ctx, id, partialView)

Patch

Update a device layout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialView** | [**PartialView**](PartialView.md)| PartialView | 

### Return type

[**View**](View.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewControllerPost

> View ViewControllerPost(ctx, view)

Post

Create a device layout

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**view** | [**View**](View.md)| View | 

### Return type

[**View**](View.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

