# \LookerApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LookerControllerGetDashboards**](LookerApi.md#LookerControllerGetDashboards) | **Get** /looker/dashboards | Get dashboards
[**LookerControllerGetLooks**](LookerApi.md#LookerControllerGetLooks) | **Get** /looker/looks | Get looks
[**LookerControllerSignLookerEmbedUrl**](LookerApi.md#LookerControllerSignLookerEmbedUrl) | **Post** /looker/sign | Sign looker embed url



## LookerControllerGetDashboards

> GetLooksResponse LookerControllerGetDashboards(ctx, )

Get dashboards

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetLooksResponse**](GetLooksResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookerControllerGetLooks

> GetLooksResponse LookerControllerGetLooks(ctx, )

Get looks

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetLooksResponse**](GetLooksResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookerControllerSignLookerEmbedUrl

> SignLookerEmbedUrlResponse LookerControllerSignLookerEmbedUrl(ctx, signLookerEmbedUrlRequest)

Sign looker embed url

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**signLookerEmbedUrlRequest** | [**SignLookerEmbedUrlRequest**](SignLookerEmbedUrlRequest.md)| SignLookerEmbedUrlRequest | 

### Return type

[**SignLookerEmbedUrlResponse**](SignLookerEmbedUrlResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

