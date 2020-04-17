# \ChannelApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelControllerDelete**](ChannelApi.md#ChannelControllerDelete) | **Delete** /channels/{id} | Delete
[**ChannelControllerGetAll**](ChannelApi.md#ChannelControllerGetAll) | **Get** /channels | Get all
[**ChannelControllerGetOne**](ChannelApi.md#ChannelControllerGetOne) | **Get** /channels/{id} | Get one
[**ChannelControllerPatch**](ChannelApi.md#ChannelControllerPatch) | **Patch** /channels/{id} | Patch
[**ChannelControllerPost**](ChannelApi.md#ChannelControllerPost) | **Post** /channels | Post



## ChannelControllerDelete

> ChannelControllerDelete(ctx, id)

Delete

Delete a channel

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


## ChannelControllerGetAll

> ChannelListResponse ChannelControllerGetAll(ctx, )

Get all

List all channels

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**ChannelListResponse**](ChannelListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelControllerGetOne

> Channel ChannelControllerGetOne(ctx, id)

Get one

Get a channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Channel**](Channel.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelControllerPatch

> Channel ChannelControllerPatch(ctx, id, partialChannel)

Patch

Update a channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialChannel** | [**PartialChannel**](PartialChannel.md)| PartialChannel | 

### Return type

[**Channel**](Channel.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelControllerPost

> Channel ChannelControllerPost(ctx, channel)

Post

Create a channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | [**Channel**](Channel.md)| Channel | 

### Return type

[**Channel**](Channel.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

