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
// PortForwardingSessionRecord struct for PortForwardingSessionRecord
type PortForwardingSessionRecord struct {
	Type string `json:"type,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	OrganizationId string `json:"organizationId,omitempty"`
	Time time.Time `json:"time"`
	Message string `json:"message,omitempty"`
	Viewed bool `json:"viewed,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	StreamName string `json:"streamName,omitempty"`
	StreamType string `json:"streamType,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	NotificationEnabled bool `json:"notificationEnabled,omitempty"`
	SessionId string `json:"sessionId"`
	UserId string `json:"userId"`
	SessionCreated time.Time `json:"sessionCreated"`
}
