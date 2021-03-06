/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
import (
	"time"
)
// PartialDeviceConfigurationTemplate struct for PartialDeviceConfigurationTemplate
type PartialDeviceConfigurationTemplate struct {
	OrganizationId string `json:"organizationId,omitempty"`
	Name string `json:"name,omitempty"`
	Document DeviceConfigurationDocument `json:"document,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
