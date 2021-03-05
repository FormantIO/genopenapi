# \StreamCurrentValueApi

All URIs are relative to *https://api-dev.formant.io/v1/queries*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StreamCurrentValueControllerQuery**](StreamCurrentValueApi.md#StreamCurrentValueControllerQuery) | **Post** /stream-current-value | Query



## StreamCurrentValueControllerQuery

> StreamCurrentValueListResponse StreamCurrentValueControllerQuery(ctx, optional)

Query

Authorized clients: viewer Authorized plans: freemium, standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StreamCurrentValueControllerQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a StreamCurrentValueControllerQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopeFilter** | [**optional.Interface of ScopeFilter**](ScopeFilter.md)| ScopeFilter | 

### Return type

[**StreamCurrentValueListResponse**](StreamCurrentValueListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

