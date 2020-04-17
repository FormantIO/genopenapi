# \FreeTeleopSettingsFormApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FreeTeleopSettingsFormControllerGetConfiguration**](FreeTeleopSettingsFormApi.md#FreeTeleopSettingsFormControllerGetConfiguration) | **Get** /free-teleop-settings-form/{id} | Get configuration
[**FreeTeleopSettingsFormControllerPostConfiguration**](FreeTeleopSettingsFormApi.md#FreeTeleopSettingsFormControllerPostConfiguration) | **Post** /free-teleop-settings-form/{id} | Post configuration



## FreeTeleopSettingsFormControllerGetConfiguration

> FreeTeleopSettingsForm FreeTeleopSettingsFormControllerGetConfiguration(ctx, id)

Get configuration

Get a free teleop settings form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**FreeTeleopSettingsForm**](FreeTeleopSettingsForm.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FreeTeleopSettingsFormControllerPostConfiguration

> FreeTeleopSettingsForm FreeTeleopSettingsFormControllerPostConfiguration(ctx, id, optional)

Post configuration

Create a free teleop settings form

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***FreeTeleopSettingsFormControllerPostConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FreeTeleopSettingsFormControllerPostConfigurationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **freeTeleopSettingsForm** | [**optional.Interface of FreeTeleopSettingsForm**](FreeTeleopSettingsForm.md)| FreeTeleopSettingsForm | 

### Return type

[**FreeTeleopSettingsForm**](FreeTeleopSettingsForm.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

