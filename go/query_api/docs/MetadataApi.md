# \MetadataApi

All URIs are relative to *https://api-dev.formant.io/v1/queries*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetadataControllerListDeviceIds**](MetadataApi.md#MetadataControllerListDeviceIds) | **Post** /metadata/device-ids | List device ids
[**MetadataControllerListMetadata**](MetadataApi.md#MetadataControllerListMetadata) | **Post** /metadata | List metadata
[**MetadataControllerListStreamNames**](MetadataApi.md#MetadataControllerListStreamNames) | **Post** /metadata/stream-names | List stream names



## MetadataControllerListDeviceIds

> UuidListResponse MetadataControllerListDeviceIds(ctx, optional)

List device ids

Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetadataControllerListDeviceIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MetadataControllerListDeviceIdsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopeFilter** | [**optional.Interface of ScopeFilter**](ScopeFilter.md)| ScopeFilter | 

### Return type

[**UuidListResponse**](UuidListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataControllerListMetadata

> MetadataListResponse MetadataControllerListMetadata(ctx, optional)

List metadata

Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetadataControllerListMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MetadataControllerListMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopeFilter** | [**optional.Interface of ScopeFilter**](ScopeFilter.md)| ScopeFilter | 

### Return type

[**MetadataListResponse**](MetadataListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataControllerListStreamNames

> StringListResponse MetadataControllerListStreamNames(ctx, optional)

List stream names

Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MetadataControllerListStreamNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MetadataControllerListStreamNamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopeFilter** | [**optional.Interface of ScopeFilter**](ScopeFilter.md)| ScopeFilter | 

### Return type

[**StringListResponse**](StringListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

