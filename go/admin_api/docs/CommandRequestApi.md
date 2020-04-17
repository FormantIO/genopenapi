# \CommandRequestApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommandRequestControllerPoll**](CommandRequestApi.md#CommandRequestControllerPoll) | **Post** /command-requests/pending | Poll
[**CommandRequestControllerPost**](CommandRequestApi.md#CommandRequestControllerPost) | **Post** /command-requests | Post



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


## CommandRequestControllerPost

> CommandRequest CommandRequestControllerPost(ctx, commandRequest)

Post

Create a command request

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandRequest** | [**CommandRequest**](CommandRequest.md)| CommandRequest | 

### Return type

[**CommandRequest**](CommandRequest.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

