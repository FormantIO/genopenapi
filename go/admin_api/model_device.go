/*
 * Formant Cloud API
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
// Device struct for Device
type Device struct {
	OrganizationId string `json:"organizationId,omitempty"`
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	PublicKey string `json:"publicKey"`
	Scope *ScopeFilter `json:"scope,omitempty"`
	DesiredAgentVersion *string `json:"desiredAgentVersion,omitempty"`
	DesiredConfigurationVersion *int64 `json:"desiredConfigurationVersion,omitempty"`
	TemporaryConfigurationVersion *int64 `json:"temporaryConfigurationVersion,omitempty"`
	TemporaryConfigurationExpiration *time.Time `json:"temporaryConfigurationExpiration,omitempty"`
	State DeviceState `json:"state,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
