# \UserApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserControllerGetAll**](UserApi.md#UserControllerGetAll) | **Get** /users | Get all
[**UserControllerGetOne**](UserApi.md#UserControllerGetOne) | **Get** /users/{id} | Get one
[**UserControllerPatch**](UserApi.md#UserControllerPatch) | **Patch** /users/{id} | Patch
[**UserControllerPost**](UserApi.md#UserControllerPost) | **Post** /users | Post



## UserControllerGetAll

> UserListResponse UserControllerGetAll(ctx, )

Get all

List all users Authorized clients: viewer Authorized plans: freemium, commercial, enterprise

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UserListResponse**](UserListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserControllerGetOne

> User UserControllerGetOne(ctx, id)

Get one

Get a user Authorized clients: viewer Authorized plans: freemium, commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserControllerPatch

> User UserControllerPatch(ctx, id, partialUser)

Patch

Update a user Authorized clients: viewer Authorized plans: freemium, commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialUser** | [**PartialUser**](PartialUser.md)| PartialUser | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserControllerPost

> User UserControllerPost(ctx, user)

Post

Create a user Authorized clients: administrator Authorized plans: freemium, commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | [**User**](User.md)| User | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

