/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// DeviceStreamDirectoryWatchConfiguration struct for DeviceStreamDirectoryWatchConfiguration
type DeviceStreamDirectoryWatchConfiguration struct {
	Type string `json:"type"`
	Directory string `json:"directory"`
	Extension string `json:"extension,omitempty"`
	FileType string `json:"fileType,omitempty"`
}
