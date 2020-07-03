# \CommandRequestApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandRequestControllerPoll**](CommandRequestApi.md#CommandRequestControllerPoll) | **Post** /command-requests/pending | Poll



## CommandRequestControllerPoll

> CommandRequestListResponse CommandRequestControllerPoll(ctx, pollCommandRequest)

Poll

Poll for command requests

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pollCommandRequest** | [**PollCommandRequest**](PollCommandRequest.md)| PollCommandRequest | 

### Return type

[**CommandRequestListResponse**](CommandRequestListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

