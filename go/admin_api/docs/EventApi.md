# \EventApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventControllerAnnotationTemplates**](EventApi.md#EventControllerAnnotationTemplates) | **Post** /events/annotation-templates | Annotation templates
[**EventControllerCount**](EventApi.md#EventControllerCount) | **Post** /events/count | Count
[**EventControllerCounts**](EventApi.md#EventControllerCounts) | **Post** /events/counts | Counts
[**EventControllerDevices**](EventApi.md#EventControllerDevices) | **Post** /events/devices | Devices
[**EventControllerEventTypes**](EventApi.md#EventControllerEventTypes) | **Post** /events/event-types | Event types
[**EventControllerExportSheet**](EventApi.md#EventControllerExportSheet) | **Post** /events/export-sheet | Export sheet
[**EventControllerGetOne**](EventApi.md#EventControllerGetOne) | **Get** /events/{id} | Get one
[**EventControllerHistogram**](EventApi.md#EventControllerHistogram) | **Post** /events/histogram | Histogram
[**EventControllerQuery**](EventApi.md#EventControllerQuery) | **Post** /events/query | Query
[**EventControllerSeek**](EventApi.md#EventControllerSeek) | **Post** /events/seek | Seek
[**EventControllerSeverities**](EventApi.md#EventControllerSeverities) | **Post** /events/severities | Severities
[**EventControllerStreamNames**](EventApi.md#EventControllerStreamNames) | **Post** /events/stream-names | Stream names
[**EventControllerStreamTypes**](EventApi.md#EventControllerStreamTypes) | **Post** /events/stream-types | Stream types
[**EventControllerTags**](EventApi.md#EventControllerTags) | **Post** /events/tags | Tags
[**EventControllerUsers**](EventApi.md#EventControllerUsers) | **Post** /events/users | Users
[**EventControllerView**](EventApi.md#EventControllerView) | **Post** /events/view | View



## EventControllerAnnotationTemplates

> UuidListResponse EventControllerAnnotationTemplates(ctx, optional)

Annotation templates

Distinct values of event annotation templates Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerAnnotationTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerAnnotationTemplatesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**UuidListResponse**](UuidListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerCount

> EventControllerCount(ctx, optional)

Count

Count events Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerCounts

> EventCounts EventControllerCounts(ctx, optional)

Counts

Count events Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerCountsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerCountsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**EventCounts**](EventCounts.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerDevices

> UuidListResponse EventControllerDevices(ctx, optional)

Devices

Distinct values of event devices Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**UuidListResponse**](UuidListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerEventTypes

> StringListResponse EventControllerEventTypes(ctx, optional)

Event types

Distinct values of event types Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerEventTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerEventTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**StringListResponse**](StringListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerExportSheet

> EventExportSheetResult EventControllerExportSheet(ctx, eventExportSheetRequest)

Export sheet

Export events as Google Sheet Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventExportSheetRequest** | [**EventExportSheetRequest**](EventExportSheetRequest.md)| EventExportSheetRequest | 

### Return type

[**EventExportSheetResult**](EventExportSheetResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerGetOne

> BaseEvent EventControllerGetOne(ctx, id)

Get one

Get an event Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**BaseEvent**](BaseEvent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerHistogram

> EventHistogram EventControllerHistogram(ctx, intervalEventFilter)

Histogram

Event histogram Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**intervalEventFilter** | [**IntervalEventFilter**](IntervalEventFilter.md)| IntervalEventFilter | 

### Return type

[**EventHistogram**](EventHistogram.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerQuery

> EventListResponse EventControllerQuery(ctx, optional)

Query

Query events Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerQueryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerQueryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventQuery** | [**optional.Interface of EventQuery**](EventQuery.md)| EventQuery | 

### Return type

[**EventListResponse**](EventListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerSeek

> EventControllerSeek(ctx, eventSeekQuery)

Seek

Seek event Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSeekQuery** | [**EventSeekQuery**](EventSeekQuery.md)| EventSeekQuery | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerSeverities

> StringListResponse EventControllerSeverities(ctx, optional)

Severities

Distinct values of event severities Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerSeveritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerSeveritiesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**StringListResponse**](StringListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerStreamNames

> StringListResponse EventControllerStreamNames(ctx, optional)

Stream names

Distinct values of event stream names Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerStreamNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerStreamNamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**StringListResponse**](StringListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerStreamTypes

> StringListResponse EventControllerStreamTypes(ctx, optional)

Stream types

Distinct values of event stream types Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerStreamTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerStreamTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**StringListResponse**](StringListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerTags

> TagsResponse EventControllerTags(ctx, optional)

Tags

Distinct values of event tags Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**TagsResponse**](TagsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerUsers

> UuidListResponse EventControllerUsers(ctx, optional)

Users

Distinct values of event users Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**UuidListResponse**](UuidListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerView

> EventControllerView(ctx, eventFilter)

View

Mark events as read Authorized clients: viewer Authorized plans: commercial, enterprise

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventFilter** | [**EventFilter**](EventFilter.md)| EventFilter | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

