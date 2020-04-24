# \EventApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventControllerCount**](EventApi.md#EventControllerCount) | **Post** /events/count | Count
[**EventControllerCounts**](EventApi.md#EventControllerCounts) | **Post** /events/counts | Counts
[**EventControllerDistinctDevices**](EventApi.md#EventControllerDistinctDevices) | **Post** /events/devices | Distinct devices
[**EventControllerDistinctEventTypes**](EventApi.md#EventControllerDistinctEventTypes) | **Post** /events/event-types | Distinct event types
[**EventControllerDistinctStreamNames**](EventApi.md#EventControllerDistinctStreamNames) | **Post** /events/stream-names | Distinct stream names
[**EventControllerDistinctStreamTypes**](EventApi.md#EventControllerDistinctStreamTypes) | **Post** /events/stream-types | Distinct stream types
[**EventControllerDistinctTags**](EventApi.md#EventControllerDistinctTags) | **Post** /events/tags | Distinct tags
[**EventControllerDistinctUsers**](EventApi.md#EventControllerDistinctUsers) | **Post** /events/users | Distinct users
[**EventControllerGetOne**](EventApi.md#EventControllerGetOne) | **Get** /events/{id} | Get one
[**EventControllerHistogram**](EventApi.md#EventControllerHistogram) | **Post** /events/histogram | Histogram
[**EventControllerQuery**](EventApi.md#EventControllerQuery) | **Post** /events/query | Query
[**EventControllerSeek**](EventApi.md#EventControllerSeek) | **Post** /events/seek | Seek
[**EventControllerView**](EventApi.md#EventControllerView) | **Post** /events/view | View



## EventControllerCount

> EventControllerCount(ctx, optional)

Count

Count events

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

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerCounts

> EventCounts EventControllerCounts(ctx, optional)

Counts

Count events

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

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerDistinctDevices

> UuidListResponse EventControllerDistinctDevices(ctx, optional)

Distinct devices

Distinct values of event devices

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerDistinctDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerDistinctDevicesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**UuidListResponse**](UuidListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerDistinctEventTypes

> EventDistinctEventTypes EventControllerDistinctEventTypes(ctx, optional)

Distinct event types

Distinct values of event types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerDistinctEventTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerDistinctEventTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**EventDistinctEventTypes**](EventDistinctEventTypes.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerDistinctStreamNames

> StringListResponse EventControllerDistinctStreamNames(ctx, optional)

Distinct stream names

Distinct values of event stream names

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerDistinctStreamNamesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerDistinctStreamNamesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**StringListResponse**](StringListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerDistinctStreamTypes

> EventDistinctStreamTypes EventControllerDistinctStreamTypes(ctx, optional)

Distinct stream types

Distinct values of event stream types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerDistinctStreamTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerDistinctStreamTypesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**EventDistinctStreamTypes**](EventDistinctStreamTypes.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerDistinctTags

> EventDistinctTags EventControllerDistinctTags(ctx, optional)

Distinct tags

Distinct values of event tags

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerDistinctTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerDistinctTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**EventDistinctTags**](EventDistinctTags.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerDistinctUsers

> UuidListResponse EventControllerDistinctUsers(ctx, optional)

Distinct users

Distinct values of event users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventControllerDistinctUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventControllerDistinctUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventFilter** | [**optional.Interface of EventFilter**](EventFilter.md)| EventFilter | 

### Return type

[**UuidListResponse**](UuidListResponse.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerGetOne

> BaseEvent EventControllerGetOne(ctx, id)

Get one

Get an event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**BaseEvent**](BaseEvent.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerHistogram

> EventHistogram EventControllerHistogram(ctx, intervalEventFilter)

Histogram

Event histogram

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**intervalEventFilter** | [**IntervalEventFilter**](IntervalEventFilter.md)| IntervalEventFilter | 

### Return type

[**EventHistogram**](EventHistogram.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerQuery

> EventListResponse EventControllerQuery(ctx, optional)

Query

Query events

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

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerSeek

> EventControllerSeek(ctx, eventSeekQuery)

Seek

Seek event

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSeekQuery** | [**EventSeekQuery**](EventSeekQuery.md)| EventSeekQuery | 

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventControllerView

> EventControllerView(ctx, eventFilter)

View

Mark events as read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventFilter** | [**EventFilter**](EventFilter.md)| EventFilter | 

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

