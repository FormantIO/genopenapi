/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// DeviceTeleopCustomStreamConfiguration struct for DeviceTeleopCustomStreamConfiguration
type DeviceTeleopCustomStreamConfiguration struct {
	Name string `json:"name"`
	RtcStreamType string `json:"rtcStreamType"`
	Mode string `json:"mode"`
	Labels []string `json:"labels,omitempty"`
	EncodeVideo bool `json:"encodeVideo,omitempty"`
}
