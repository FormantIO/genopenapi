/*
 * Formant Cloud API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// DeviceStreamConfiguration struct for DeviceStreamConfiguration
type DeviceStreamConfiguration struct {
	Name string `json:"name"`
	Tags map[string]string `json:"tags,omitempty"`
	Configuration interface{} `json:"configuration"`
	ThrottleHz *float32 `json:"throttleHz,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
}
