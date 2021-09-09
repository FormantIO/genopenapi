/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// DeviceTeleopHardwareStreamConfiguration struct for DeviceTeleopHardwareStreamConfiguration
type DeviceTeleopHardwareStreamConfiguration struct {
	Name string `json:"name"`
	RtcStreamType string `json:"rtcStreamType"`
	Mode string `json:"mode"`
	HwDescriptor string `json:"hwDescriptor"`
	HardwareType string `json:"hardwareType"`
	RtspEncodingNeeded bool `json:"rtspEncodingNeeded,omitempty"`
	IsOnvif bool `json:"isOnvif,omitempty"`
	IpCamUsername string `json:"ipCamUsername,omitempty"`
	IpCamPassword string `json:"ipCamPassword,omitempty"`
	Quality string `json:"quality"`
}
