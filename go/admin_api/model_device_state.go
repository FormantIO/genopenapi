/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// DeviceState struct for DeviceState
type DeviceState struct {
	AgentVersion *string `json:"agentVersion,omitempty"`
	ReportedConfiguration *DeviceReportedConfigurationState `json:"reportedConfiguration,omitempty"`
	HwInfo *HwInfo `json:"hwInfo,omitempty"`
	Ros *DeviceRosState `json:"ros,omitempty"`
	Env *map[string]string `json:"env,omitempty"`
	OtaEnabled *bool `json:"otaEnabled,omitempty"`
	OnDemand *OnDemandState `json:"onDemand,omitempty"`
	CommandProgress *[]interface{} `json:"commandProgress,omitempty"`
	Version *string `json:"version,omitempty"`
}
