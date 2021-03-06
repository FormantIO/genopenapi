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
// TriggeredEvent struct for TriggeredEvent
type TriggeredEvent struct {
	Type string `json:"type,omitempty"`
	Severity string `json:"severity"`
	EventTriggerId string `json:"eventTriggerId"`
	Interval int64 `json:"interval"`
	IntervalStart time.Time `json:"intervalStart,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	OrganizationId string `json:"organizationId,omitempty"`
	Time time.Time `json:"time"`
	EndTime *time.Time `json:"endTime,omitempty"`
	Message string `json:"message,omitempty"`
	Viewed bool `json:"viewed,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	StreamName string `json:"streamName,omitempty"`
	StreamType string `json:"streamType,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	NotificationEnabled bool `json:"notificationEnabled,omitempty"`
}
