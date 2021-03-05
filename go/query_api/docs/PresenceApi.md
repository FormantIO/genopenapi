# \PresenceApi

All URIs are relative to *https://api-dev.formant.io/v1/queries*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PresenceControllerCount**](PresenceApi.md#PresenceControllerCount) | **Post** /presence | Count



## PresenceControllerCount

> IsoDateListResponse PresenceControllerCount(ctx, intervalQuery)

Count

Authorized clients: viewer Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**intervalQuery** | [**IntervalQuery**](IntervalQuery.md)| IntervalQuery | 

### Return type

[**IsoDateListResponse**](IsoDateListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

