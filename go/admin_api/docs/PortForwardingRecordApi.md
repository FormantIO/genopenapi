# \PortForwardingRecordApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortForwardingRecordControllerPost**](PortForwardingRecordApi.md#PortForwardingRecordControllerPost) | **Post** /port-forwarding-session-records | Post



## PortForwardingRecordControllerPost

> PortForwardingSessionRecord PortForwardingRecordControllerPost(ctx, portForwardingSessionRecord)

Post

Create a port forwarding session record.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portForwardingSessionRecord** | [**PortForwardingSessionRecord**](PortForwardingSessionRecord.md)| PortForwardingSessionRecord | 

### Return type

[**PortForwardingSessionRecord**](PortForwardingSessionRecord.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

