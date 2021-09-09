# \CommandTemplateApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandTemplateControllerDelete**](CommandTemplateApi.md#CommandTemplateControllerDelete) | **Delete** /command-templates/{id} | Delete
[**CommandTemplateControllerGetOne**](CommandTemplateApi.md#CommandTemplateControllerGetOne) | **Get** /command-templates/{id} | Get one
[**CommandTemplateControllerList**](CommandTemplateApi.md#CommandTemplateControllerList) | **Get** /command-templates/ | List
[**CommandTemplateControllerPatch**](CommandTemplateApi.md#CommandTemplateControllerPatch) | **Patch** /command-templates/{id} | Patch
[**CommandTemplateControllerPost**](CommandTemplateApi.md#CommandTemplateControllerPost) | **Post** /command-templates | Post



## CommandTemplateControllerDelete

> CommandTemplateControllerDelete(ctx, id)

Delete

Delete a command template Authorized clients: administrator Authorized plans: starter, premium

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


## CommandTemplateControllerGetOne

> CommandTemplate CommandTemplateControllerGetOne(ctx, id)

Get one

Get a command template Authorized clients: viewer Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**CommandTemplate**](CommandTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommandTemplateControllerList

> CommandTemplateListResponse CommandTemplateControllerList(ctx, )

List

List command templates Authorized clients: viewer Authorized plans: starter, premium

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**CommandTemplateListResponse**](CommandTemplateListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommandTemplateControllerPatch

> CommandTemplate CommandTemplateControllerPatch(ctx, id, partialCommandTemplate)

Patch

Update a command template Authorized clients: administrator Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialCommandTemplate** | [**PartialCommandTemplate**](PartialCommandTemplate.md)| PartialCommandTemplate | 

### Return type

[**CommandTemplate**](CommandTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommandTemplateControllerPost

> CommandTemplate CommandTemplateControllerPost(ctx, commandTemplate)

Post

Create a command template Authorized clients: administrator Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandTemplate** | [**CommandTemplate**](CommandTemplate.md)| CommandTemplate | 

### Return type

[**CommandTemplate**](CommandTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

