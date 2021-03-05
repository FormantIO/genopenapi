# \QueryApi

All URIs are relative to *https://api-dev.formant.io/v1/queries*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryControllerExportSheet**](QueryApi.md#QueryControllerExportSheet) | **Post** /queries/export-sheet | Export sheet
[**QueryControllerQuery**](QueryApi.md#QueryControllerQuery) | **Post** /queries | Query



## QueryControllerExportSheet

> TelemetryExportSheetResult QueryControllerExportSheet(ctx, telemetryExportSheetRequest)

Export sheet

Authorized clients: operator Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**telemetryExportSheetRequest** | [**TelemetryExportSheetRequest**](TelemetryExportSheetRequest.md)| TelemetryExportSheetRequest | 

### Return type

[**TelemetryExportSheetResult**](TelemetryExportSheetResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryControllerQuery

> StreamDataListResponse QueryControllerQuery(ctx, query)

Query

Authorized clients: viewer Authorized plans: standard, premium, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | [**Query**](Query.md)| Query | 

### Return type

[**StreamDataListResponse**](StreamDataListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

