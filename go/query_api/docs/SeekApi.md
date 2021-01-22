# \SeekApi

All URIs are relative to *https://api-dev.formant.io/v1/queries*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SeekControllerSeek**](SeekApi.md#SeekControllerSeek) | **Post** /seek | Seek



## SeekControllerSeek

> SeekResult SeekControllerSeek(ctx, seekQuery)

Seek

Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seekQuery** | [**SeekQuery**](SeekQuery.md)| SeekQuery | 

### Return type

[**SeekResult**](SeekResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

