# PartialDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Type** | **string** |  | [optional] 
**Tags** | **map[string]string** |  | [optional] 
**PublicKey** | **string** |  | [optional] 
**Scope** | Pointer to [**ScopeFilter**](ScopeFilter.md) |  | [optional] 
**DesiredAgentVersion** | Pointer to **string** |  | [optional] 
**DesiredConfigurationVersion** | Pointer to **int64** |  | [optional] 
**TemporaryConfigurationVersion** | Pointer to **int64** |  | [optional] 
**TemporaryConfigurationExpiration** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**TemporaryConfigurationTemplateId** | Pointer to **string** |  | [optional] 
**Followers** | [**[]DeviceFollower**](DeviceFollower.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**State** | [**DeviceState**](DeviceState.md) |  | [optional] 
**Enabled** | **bool** |  | [optional] 
**DisabledAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Id** | **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


