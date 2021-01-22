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
// TagTemplate struct for TagTemplate
type TagTemplate struct {
	OrganizationId string `json:"organizationId,omitempty"`
	TagKey interface{} `json:"tagKey"`
	IsFleet bool `json:"isFleet,omitempty"`
	IsTelemetryFilter bool `json:"isTelemetryFilter,omitempty"`
	IsEventFilter bool `json:"isEventFilter,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
