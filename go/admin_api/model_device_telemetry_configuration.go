/*
 * Formant Cloud API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// DeviceTelemetryConfiguration struct for DeviceTelemetryConfiguration
type DeviceTelemetryConfiguration struct {
	Streams []DeviceStreamConfiguration `json:"streams,omitempty"`
	Ros DeviceRosConfiguration `json:"ros,omitempty"`
}
