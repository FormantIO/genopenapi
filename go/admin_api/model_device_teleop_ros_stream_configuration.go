/*
 * Formant Cloud API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// DeviceTeleopRosStreamConfiguration struct for DeviceTeleopRosStreamConfiguration
type DeviceTeleopRosStreamConfiguration struct {
	TopicName string `json:"topicName"`
	TopicType string `json:"topicType"`
	Mode string `json:"mode"`
	EncodeVideo bool `json:"encodeVideo,omitempty"`
}
