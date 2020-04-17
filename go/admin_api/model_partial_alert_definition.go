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
// PartialAlertDefinition struct for PartialAlertDefinition
type PartialAlertDefinition struct {
	OrganizationId string `json:"organizationId,omitempty"`
	Message string `json:"message,omitempty"`
	Stream string `json:"stream,omitempty"`
	Condition interface{} `json:"condition,omitempty"`
	Interval int64 `json:"interval,omitempty"`
	Severity string `json:"severity,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	NotificationEnabled bool `json:"notificationEnabled,omitempty"`
	Commands []AlertCommand `json:"commands,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
