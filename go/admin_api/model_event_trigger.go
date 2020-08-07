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
// EventTrigger struct for EventTrigger
type EventTrigger struct {
	OrganizationId string `json:"organizationId,omitempty"`
	Message string `json:"message"`
	Stream string `json:"stream"`
	Condition interface{} `json:"condition"`
	Interval int64 `json:"interval"`
	Severity string `json:"severity"`
	Enabled bool `json:"enabled,omitempty"`
	TriggeredConfiguration *TriggeredConfiguration `json:"triggeredConfiguration"`
	Tags map[string]string `json:"tags"`
	NotificationEnabled bool `json:"notificationEnabled"`
	Commands []EventTriggerCommand `json:"commands"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
