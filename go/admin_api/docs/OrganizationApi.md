# \OrganizationApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationControllerGetOne**](OrganizationApi.md#OrganizationControllerGetOne) | **Get** /organizations/{id} | Get one
[**OrganizationControllerGetSlackInfo**](OrganizationApi.md#OrganizationControllerGetSlackInfo) | **Post** /organizations/slack-info | Get slack info
[**OrganizationControllerPatch**](OrganizationApi.md#OrganizationControllerPatch) | **Patch** /organizations/{id} | Patch



## OrganizationControllerGetOne

> Organization OrganizationControllerGetOne(ctx, id)

Get one

Get an organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**Organization**](Organization.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationControllerGetSlackInfo

> OrganizationControllerGetSlackInfo(ctx, )

Get slack info

Private API provides Slack integration configuration

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationControllerPatch

> Organization OrganizationControllerPatch(ctx, id, partialOrganization)

Patch

Update an organization

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**partialOrganization** | [**PartialOrganization**](PartialOrganization.md)| PartialOrganization | 

### Return type

[**Organization**](Organization.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

