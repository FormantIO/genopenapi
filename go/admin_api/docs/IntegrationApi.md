# \IntegrationApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationControllerGoogle**](IntegrationApi.md#IntegrationControllerGoogle) | **Post** /integrations/google/auth | Google
[**IntegrationControllerSlack**](IntegrationApi.md#IntegrationControllerSlack) | **Post** /integrations/slack/auth | Slack



## IntegrationControllerGoogle

> GoogleInfo IntegrationControllerGoogle(ctx, googleAuthRequest)

Google

Create a Google integration Authorized clients: administrator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**googleAuthRequest** | [**GoogleAuthRequest**](GoogleAuthRequest.md)| GoogleAuthRequest | 

### Return type

[**GoogleInfo**](GoogleInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationControllerSlack

> SlackInfo IntegrationControllerSlack(ctx, slackAuthRequest)

Slack

Create a Slack integration Authorized clients: administrator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slackAuthRequest** | [**SlackAuthRequest**](SlackAuthRequest.md)| SlackAuthRequest | 

### Return type

[**SlackInfo**](SlackInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

