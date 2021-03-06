/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// DeviceTeleopConfiguration struct for DeviceTeleopConfiguration
type DeviceTeleopConfiguration struct {
	RosStreams []DeviceTeleopRosStreamConfiguration `json:"rosStreams,omitempty"`
	CustomStreams []DeviceTeleopCustomStreamConfiguration `json:"customStreams,omitempty"`
	HardwareStreams []DeviceTeleopHardwareStreamConfiguration `json:"hardwareStreams,omitempty"`
	Joysticks []interface{} `json:"joysticks,omitempty"`
	Views []interface{} `json:"views,omitempty"`
	ArmSwitch bool `json:"armSwitch,omitempty"`
}
