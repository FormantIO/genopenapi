# \AlertDefinitionApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertDefinitionControllerGetAll**](AlertDefinitionApi.md#AlertDefinitionControllerGetAll) | **Get** /alert-definitions | Get all
[**AlertDefinitionControllerGetOne**](AlertDefinitionApi.md#AlertDefinitionControllerGetOne) | **Get** /alert-definitions/{id} | Get one
[**AlertDefinitionControllerPatch**](AlertDefinitionApi.md#AlertDefinitionControllerPatch) | **Patch** /alert-definitions/{id} | Patch
[**AlertDefinitionControllerPost**](AlertDefinitionApi.md#AlertDefinitionControllerPost) | **Post** /alert-definitions | Post



## AlertDefinitionControllerGetAll

> AlertDefinitionListResponse AlertDefinitionControllerGetAll(ctx, )

Get all

List all alert definitions

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AlertDefinitionListResponse**](AlertDefinitionListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionControllerGetOne

> AlertDefinition AlertDefinitionControllerGetOne(ctx, id)

Get one

Get an alert definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionControllerPatch

> AlertDefinition AlertDefinitionControllerPatch(ctx, id, partialAlertDefinition)

Patch

Update alert definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialAlertDefinition** | [**PartialAlertDefinition**](PartialAlertDefinition.md)| PartialAlertDefinition | 

### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertDefinitionControllerPost

> AlertDefinition AlertDefinitionControllerPost(ctx, alertDefinition)

Post

Create alert definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertDefinition** | [**AlertDefinition**](AlertDefinition.md)| AlertDefinition | 

### Return type

[**AlertDefinition**](AlertDefinition.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

