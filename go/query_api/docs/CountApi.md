# \CountApi

All URIs are relative to *https://api.formant.io/v1/queries*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountControllerHistory**](CountApi.md#CountControllerHistory) | **Post** /counts/history | History



## CountControllerHistory

> CountHistory CountControllerHistory(ctx, countHistoryQuery)

History

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countHistoryQuery** | [**CountHistoryQuery**](CountHistoryQuery.md)| CountHistoryQuery | 

### Return type

[**CountHistory**](CountHistory.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

