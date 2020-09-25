# \DeviceApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceControllerCount**](DeviceApi.md#DeviceControllerCount) | **Post** /devices/count | Count
[**DeviceControllerFilter**](DeviceApi.md#DeviceControllerFilter) | **Post** /devices/filter | Filter
[**DeviceControllerGenerateDeviceProvisioningToken**](DeviceApi.md#DeviceControllerGenerateDeviceProvisioningToken) | **Post** /devices/{id}/provisioning-token | Generate device provisioning token
[**DeviceControllerGetConfiguration**](DeviceApi.md#DeviceControllerGetConfiguration) | **Get** /devices/{id}/configurations/{version} | Get configuration
[**DeviceControllerGetOne**](DeviceApi.md#DeviceControllerGetOne) | **Get** /devices/{id} | Get one
[**DeviceControllerGetTags**](DeviceApi.md#DeviceControllerGetTags) | **Post** /devices/tags | Get tags
[**DeviceControllerGetUpdatedAgentVersion**](DeviceApi.md#DeviceControllerGetUpdatedAgentVersion) | **Get** /devices/{id}/updated-agent-version | Get updated agent version
[**DeviceControllerGetUpdatedConfiguration**](DeviceApi.md#DeviceControllerGetUpdatedConfiguration) | **Get** /devices/{id}/updated-configuration | Get updated configuration
[**DeviceControllerPatch**](DeviceApi.md#DeviceControllerPatch) | **Patch** /devices/{id} | Patch
[**DeviceControllerPost**](DeviceApi.md#DeviceControllerPost) | **Post** /devices | Post
[**DeviceControllerPostConfiguration**](DeviceApi.md#DeviceControllerPostConfiguration) | **Post** /devices/{id}/configurations | Post configuration
[**DeviceControllerProvisionDevice**](DeviceApi.md#DeviceControllerProvisionDevice) | **Post** /devices/provision | Provision device
[**DeviceControllerQuery**](DeviceApi.md#DeviceControllerQuery) | **Post** /devices/query | Query
[**DeviceControllerUnprovisionDevice**](DeviceApi.md#DeviceControllerUnprovisionDevice) | **Post** /devices/{id}/unprovision | Unprovision device
[**DeviceControllerValidateStreamConfiguration**](DeviceApi.md#DeviceControllerValidateStreamConfiguration) | **Post** /devices/validate-stream-configuration | Validate stream configuration



## DeviceControllerCount

> DeviceControllerCount(ctx, optional)

Count

Count devices Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceControllerCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceControllerCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceQuery** | [**optional.Interface of DeviceQuery**](DeviceQuery.md)| DeviceQuery | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerFilter

> DeviceListResponse DeviceControllerFilter(ctx, optional)

Filter

Query devices by name and/or tags Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceControllerFilterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceControllerFilterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceQuery** | [**optional.Interface of DeviceQuery**](DeviceQuery.md)| DeviceQuery | 

### Return type

[**DeviceListResponse**](DeviceListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerGenerateDeviceProvisioningToken

> DeviceProvisioning DeviceControllerGenerateDeviceProvisioningToken(ctx, id)

Generate device provisioning token

Generate a device provisioning token Authorized clients: administrator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**DeviceProvisioning**](DeviceProvisioning.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerGetConfiguration

> DeviceConfiguration DeviceControllerGetConfiguration(ctx, id, version)

Get configuration

Get a device configuration Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**version** | **string**|  | 

### Return type

[**DeviceConfiguration**](DeviceConfiguration.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerGetOne

> Device DeviceControllerGetOne(ctx, id)

Get one

Get a device Authorized clients: viewer, device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerGetTags

> TagsResponse DeviceControllerGetTags(ctx, optional)

Get tags

Get tags across all devices Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceControllerGetTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceControllerGetTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceQuery** | [**optional.Interface of DeviceQuery**](DeviceQuery.md)| DeviceQuery | 

### Return type

[**TagsResponse**](TagsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerGetUpdatedAgentVersion

> UpdatedAgentVersionResponse DeviceControllerGetUpdatedAgentVersion(ctx, id, optional)

Get updated agent version

Check for agent version updates Authorized clients: device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***DeviceControllerGetUpdatedAgentVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceControllerGetUpdatedAgentVersionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportedAgentVersion** | **optional.String**|  | 

### Return type

[**UpdatedAgentVersionResponse**](UpdatedAgentVersionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerGetUpdatedConfiguration

> UpdatedConfigurationResponse DeviceControllerGetUpdatedConfiguration(ctx, id, optional)

Get updated configuration

Check for updated device configuration Authorized clients: device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***DeviceControllerGetUpdatedConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceControllerGetUpdatedConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appVersion** | **optional.String**|  | 
 **reportedConfigurationVersion** | **optional.Int64**|  | 

### Return type

[**UpdatedConfigurationResponse**](UpdatedConfigurationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerPatch

> Device DeviceControllerPatch(ctx, id, partialDevice)

Patch

Update a device Authorized clients: administrator, device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialDevice** | [**PartialDevice**](PartialDevice.md)| PartialDevice | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerPost

> Device DeviceControllerPost(ctx, device)

Post

Create a device Authorized clients: administrator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**device** | [**Device**](Device.md)| Device | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerPostConfiguration

> DeviceConfiguration DeviceControllerPostConfiguration(ctx, id, deviceConfiguration)

Post configuration

Create a device configuration Authorized clients: administrator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**deviceConfiguration** | [**DeviceConfiguration**](DeviceConfiguration.md)| DeviceConfiguration | 

### Return type

[**DeviceConfiguration**](DeviceConfiguration.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerProvisionDevice

> Device DeviceControllerProvisionDevice(ctx, deviceProvisioningRequest)

Provision device

Provision a device

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceProvisioningRequest** | [**DeviceProvisioningRequest**](DeviceProvisioningRequest.md)| DeviceProvisioningRequest | 

### Return type

[**Device**](Device.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerQuery

> DeviceListResponse DeviceControllerQuery(ctx, optional)

Query

Query devices by name and/or tags Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeviceControllerQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeviceControllerQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceQuery** | [**optional.Interface of DeviceQuery**](DeviceQuery.md)| DeviceQuery | 

### Return type

[**DeviceListResponse**](DeviceListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceControllerUnprovisionDevice

> DeviceControllerUnprovisionDevice(ctx, id)

Unprovision device

Unprovision a device Authorized clients: administrator

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


## DeviceControllerValidateStreamConfiguration

> DeviceControllerValidateStreamConfiguration(ctx, deviceStreamConfiguration)

Validate stream configuration

Validate a device stream configuration Authorized clients: administrator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceStreamConfiguration** | [**DeviceStreamConfiguration**](DeviceStreamConfiguration.md)| DeviceStreamConfiguration | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

