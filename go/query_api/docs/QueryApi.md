# \QueryApi

All URIs are relative to *https://api.formant.io/v1/queries*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryControllerQuery**](QueryApi.md#QueryControllerQuery) | **Post** /queries | Query



## QueryControllerQuery

> StreamDataListResponse QueryControllerQuery(ctx, query)

Query

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | [**Query**](Query.md)| Query | 

### Return type

[**StreamDataListResponse**](StreamDataListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

