# \FileApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileControllerBeginUpload**](FileApi.md#FileControllerBeginUpload) | **Post** /files/begin-upload | Begin upload
[**FileControllerCompleteUpload**](FileApi.md#FileControllerCompleteUpload) | **Post** /files/complete-upload | Complete upload
[**FileControllerQuery**](FileApi.md#FileControllerQuery) | **Post** /files/query | Query



## FileControllerBeginUpload

> BeginUploadResponse FileControllerBeginUpload(ctx, beginUploadRequest)

Begin upload

Initiate a file upload Authorized clients: administrator Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beginUploadRequest** | [**BeginUploadRequest**](BeginUploadRequest.md)| BeginUploadRequest | 

### Return type

[**BeginUploadResponse**](BeginUploadResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileControllerCompleteUpload

> FileControllerCompleteUpload(ctx, completeUploadRequest)

Complete upload

Complete a file upload Authorized clients: administrator Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**completeUploadRequest** | [**CompleteUploadRequest**](CompleteUploadRequest.md)| CompleteUploadRequest | 

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


## FileControllerQuery

> QueryFilesResponse FileControllerQuery(ctx, queryFilesRequest)

Query

Query files Authorized clients: operator, device Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queryFilesRequest** | [**QueryFilesRequest**](QueryFilesRequest.md)| QueryFilesRequest | 

### Return type

[**QueryFilesResponse**](QueryFilesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

