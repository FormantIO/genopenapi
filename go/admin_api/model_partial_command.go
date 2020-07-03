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
// PartialCommand struct for PartialCommand
type PartialCommand struct {
	OrganizationId string `json:"organizationId,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	UserId string `json:"userId,omitempty"`
	EventTriggerId string `json:"eventTriggerId,omitempty"`
	CommandTemplateId string `json:"commandTemplateId,omitempty"`
	Command string `json:"command,omitempty"`
	Parameter CommandParameter `json:"parameter,omitempty"`
	DeliveredAt time.Time `json:"deliveredAt,omitempty"`
	CanceledAt time.Time `json:"canceledAt,omitempty"`
	RespondedAt time.Time `json:"respondedAt,omitempty"`
	Success bool `json:"success,omitempty"`
	StreamName string `json:"streamName,omitempty"`
	StreamType string `json:"streamType,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}