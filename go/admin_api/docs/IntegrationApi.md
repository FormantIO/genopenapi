# \IntegrationApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationControllerPost**](IntegrationApi.md#IntegrationControllerPost) | **Post** /integrations/slack/auth | Post



## IntegrationControllerPost

> SlackInfo IntegrationControllerPost(ctx, slackAuthRequest)

Post

Create a slack integration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slackAuthRequest** | [**SlackAuthRequest**](SlackAuthRequest.md)| SlackAuthRequest | 

### Return type

[**SlackInfo**](SlackInfo.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

