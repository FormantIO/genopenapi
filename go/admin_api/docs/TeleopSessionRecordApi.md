# \TeleopSessionRecordApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeleopSessionRecordControllerPost**](TeleopSessionRecordApi.md#TeleopSessionRecordControllerPost) | **Post** /teleop-session-records | Post



## TeleopSessionRecordControllerPost

> TeleopSessionRecord TeleopSessionRecordControllerPost(ctx, teleopSessionRecord)

Post

Create a teleoperation session record.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teleopSessionRecord** | [**TeleopSessionRecord**](TeleopSessionRecord.md)| TeleopSessionRecord | 

### Return type

[**TeleopSessionRecord**](TeleopSessionRecord.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

