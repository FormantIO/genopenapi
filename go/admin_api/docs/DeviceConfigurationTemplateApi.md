# \DeviceConfigurationTemplateApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceConfigurationTemplateControllerDelete**](DeviceConfigurationTemplateApi.md#DeviceConfigurationTemplateControllerDelete) | **Delete** /device-configuration-templates/{id} | Delete
[**DeviceConfigurationTemplateControllerGetOne**](DeviceConfigurationTemplateApi.md#DeviceConfigurationTemplateControllerGetOne) | **Get** /device-configuration-templates/{id} | Get one
[**DeviceConfigurationTemplateControllerList**](DeviceConfigurationTemplateApi.md#DeviceConfigurationTemplateControllerList) | **Get** /device-configuration-templates/ | List
[**DeviceConfigurationTemplateControllerPatch**](DeviceConfigurationTemplateApi.md#DeviceConfigurationTemplateControllerPatch) | **Patch** /device-configuration-templates/{id} | Patch
[**DeviceConfigurationTemplateControllerPost**](DeviceConfigurationTemplateApi.md#DeviceConfigurationTemplateControllerPost) | **Post** /device-configuration-templates | Post



## DeviceConfigurationTemplateControllerDelete

> DeviceConfigurationTemplateControllerDelete(ctx, id)

Delete

Delete a device configuration template Authorized clients: administrator Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceConfigurationTemplateControllerGetOne

> DeviceConfigurationTemplate DeviceConfigurationTemplateControllerGetOne(ctx, id)

Get one

Get a device configuration template Authorized clients: administrator Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**DeviceConfigurationTemplate**](DeviceConfigurationTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceConfigurationTemplateControllerList

> DeviceConfigurationTemplateListResponse DeviceConfigurationTemplateControllerList(ctx, )

List

List device configuration templates Authorized clients: administrator Authorized plans: standard, premium, enterprise

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DeviceConfigurationTemplateListResponse**](DeviceConfigurationTemplateListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceConfigurationTemplateControllerPatch

> DeviceConfigurationTemplate DeviceConfigurationTemplateControllerPatch(ctx, id, partialDeviceConfigurationTemplate)

Patch

Update a device configuration template Authorized clients: administrator Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialDeviceConfigurationTemplate** | [**PartialDeviceConfigurationTemplate**](PartialDeviceConfigurationTemplate.md)| PartialDeviceConfigurationTemplate | 

### Return type

[**DeviceConfigurationTemplate**](DeviceConfigurationTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceConfigurationTemplateControllerPost

> DeviceConfigurationTemplate DeviceConfigurationTemplateControllerPost(ctx, deviceConfigurationTemplate)

Post

Create a device configuration template Authorized clients: administrator Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceConfigurationTemplate** | [**DeviceConfigurationTemplate**](DeviceConfigurationTemplate.md)| DeviceConfigurationTemplate | 

### Return type

[**DeviceConfigurationTemplate**](DeviceConfigurationTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

