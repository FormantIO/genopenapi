# \AnnotationApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnnotationControllerDelete**](AnnotationApi.md#AnnotationControllerDelete) | **Delete** /annotations/{id} | Delete
[**AnnotationControllerGetOne**](AnnotationApi.md#AnnotationControllerGetOne) | **Get** /annotations/{id} | Get one
[**AnnotationControllerPatch**](AnnotationApi.md#AnnotationControllerPatch) | **Patch** /annotations/{id} | Patch
[**AnnotationControllerPost**](AnnotationApi.md#AnnotationControllerPost) | **Post** /annotations | Post



## AnnotationControllerDelete

> AnnotationControllerDelete(ctx, id)

Delete

Delete an annotation

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


## AnnotationControllerGetOne

> Annotation AnnotationControllerGetOne(ctx, id)

Get one

Get an annotation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Annotation**](Annotation.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnotationControllerPatch

> Annotation AnnotationControllerPatch(ctx, id, partialAnnotation)

Patch

Update an annotation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialAnnotation** | [**PartialAnnotation**](PartialAnnotation.md)| PartialAnnotation | 

### Return type

[**Annotation**](Annotation.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnnotationControllerPost

> Annotation AnnotationControllerPost(ctx, annotation)

Post

Create an annotation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**annotation** | [**Annotation**](Annotation.md)| Annotation | 

### Return type

[**Annotation**](Annotation.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

