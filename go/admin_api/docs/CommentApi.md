# \CommentApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommentControllerDelete**](CommentApi.md#CommentControllerDelete) | **Delete** /comments/{id} | Delete
[**CommentControllerGetOne**](CommentApi.md#CommentControllerGetOne) | **Get** /comments/{id} | Get one
[**CommentControllerPatch**](CommentApi.md#CommentControllerPatch) | **Patch** /comments/{id} | Patch
[**CommentControllerPost**](CommentApi.md#CommentControllerPost) | **Post** /comments | Post



## CommentControllerDelete

> CommentControllerDelete(ctx, id)

Delete

Delete a comment Authorized clients: operator Authorized plans: standard, premium, enterprise

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


## CommentControllerGetOne

> Comment CommentControllerGetOne(ctx, id)

Get one

Get a comment Authorized clients: viewer Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Comment**](Comment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentControllerPatch

> Comment CommentControllerPatch(ctx, id, partialComment)

Patch

Update a comment Authorized clients: operator Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialComment** | [**PartialComment**](PartialComment.md)| PartialComment | 

### Return type

[**Comment**](Comment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommentControllerPost

> Comment CommentControllerPost(ctx, comment)

Post

Create a comment Authorized clients: operator Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**comment** | [**Comment**](Comment.md)| Comment | 

### Return type

[**Comment**](Comment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

