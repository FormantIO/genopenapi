# \OrganizationApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationControllerQueryUsageRecords**](OrganizationApi.md#OrganizationControllerQueryUsageRecords) | **Post** /organizations/{id}/usage-record/query | Query usage records



## OrganizationControllerQueryUsageRecords

> UsageRecordQueryResponse OrganizationControllerQueryUsageRecords(ctx, id, usageRecordQuery)

Query usage records

Query for usage records Authorized clients: administrator Authorized plans: freemium, starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**usageRecordQuery** | [**UsageRecordQuery**](UsageRecordQuery.md)| UsageRecordQuery | 

### Return type

[**UsageRecordQueryResponse**](UsageRecordQueryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

