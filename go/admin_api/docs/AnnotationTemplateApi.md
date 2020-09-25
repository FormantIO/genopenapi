# \AnnotationTemplateApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnnotationTemplateControllerDelete**](AnnotationTemplateApi.md#AnnotationTemplateControllerDelete) | **Delete** /annotation-templates/{id} | Delete
[**AnnotationTemplateControllerGetOne**](AnnotationTemplateApi.md#AnnotationTemplateControllerGetOne) | **Get** /annotation-templates/{id} | Get one
[**AnnotationTemplateControllerInspectSpreadsheet**](AnnotationTemplateApi.md#AnnotationTemplateControllerInspectSpreadsheet) | **Post** /annotation-templates/inspect-spreadsheet | Inspect spreadsheet
[**AnnotationTemplateControllerList**](AnnotationTemplateApi.md#AnnotationTemplateControllerList) | **Get** /annotation-templates/ | List
[**AnnotationTemplateControllerPatch**](AnnotationTemplateApi.md#AnnotationTemplateControllerPatch) | **Patch** /annotation-templates/{id} | Patch
[**AnnotationTemplateControllerPost**](AnnotationTemplateApi.md#AnnotationTemplateControllerPost) | **Post** /annotation-templates | Post
[**AnnotationTemplateControllerQueryAnnotationFieldValues**](AnnotationTemplateApi.md#AnnotationTemplateControllerQueryAnnotationFieldValues) | **Post** /annotation-templates/query | Query annotation field values



## AnnotationTemplateControllerDelete

> AnnotationTemplateControllerDelete(ctx, id)

Delete

Delete an annotation template Authorized clients: administrator

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


## AnnotationTemplateControllerGetOne

> AnnotationTemplate AnnotationTemplateControllerGetOne(ctx, id)

Get one

Get an annotation template Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**AnnotationTemplate**](AnnotationTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnotationTemplateControllerInspectSpreadsheet

> InspectSpreadsheetResponse AnnotationTemplateControllerInspectSpreadsheet(ctx, inspectSpreadsheetRequest)

Inspect spreadsheet

Returns information about the Spreadsheet and its content Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inspectSpreadsheetRequest** | [**InspectSpreadsheetRequest**](InspectSpreadsheetRequest.md)| InspectSpreadsheetRequest | 

### Return type

[**InspectSpreadsheetResponse**](InspectSpreadsheetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnotationTemplateControllerList

> AnnotationTemplateListResponse AnnotationTemplateControllerList(ctx, )

List

List annotation templates Authorized clients: viewer

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**AnnotationTemplateListResponse**](AnnotationTemplateListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnotationTemplateControllerPatch

> AnnotationTemplate AnnotationTemplateControllerPatch(ctx, id, partialAnnotationTemplate)

Patch

Update an annotation template Authorized clients: administrator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialAnnotationTemplate** | [**PartialAnnotationTemplate**](PartialAnnotationTemplate.md)| PartialAnnotationTemplate | 

### Return type

[**AnnotationTemplate**](AnnotationTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnotationTemplateControllerPost

> AnnotationTemplate AnnotationTemplateControllerPost(ctx, annotationTemplate)

Post

Create an annotation template Authorized clients: administrator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**annotationTemplate** | [**AnnotationTemplate**](AnnotationTemplate.md)| AnnotationTemplate | 

### Return type

[**AnnotationTemplate**](AnnotationTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnotationTemplateControllerQueryAnnotationFieldValues

> AnnotationFieldValuesResponse AnnotationTemplateControllerQueryAnnotationFieldValues(ctx, annotationFieldValuesRequest)

Query annotation field values

Query annotation template field values Authorized clients: viewer

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**annotationFieldValuesRequest** | [**AnnotationFieldValuesRequest**](AnnotationFieldValuesRequest.md)| AnnotationFieldValuesRequest | 

### Return type

[**AnnotationFieldValuesResponse**](AnnotationFieldValuesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

