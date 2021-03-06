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
// EventSeekQuery struct for EventSeekQuery
type EventSeekQuery struct {
	Direction string `json:"direction"`
	From time.Time `json:"from"`
	AgentIds []string `json:"agentIds,omitempty"`
	DeviceIds []string `json:"deviceIds,omitempty"`
	Names []string `json:"names,omitempty"`
	Types []string `json:"types,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	NotNames []string `json:"notNames,omitempty"`
	Id string `json:"id,omitempty"`
	Viewed bool `json:"viewed,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	Message string `json:"message,omitempty"`
	Start time.Time `json:"start,omitempty"`
	End time.Time `json:"end,omitempty"`
	EventTypes []string `json:"eventTypes,omitempty"`
	NotificationEnabled bool `json:"notificationEnabled,omitempty"`
	UserIds []string `json:"userIds,omitempty"`
	AnnotationTemplateIds []string `json:"annotationTemplateIds,omitempty"`
	DisableNullMatches bool `json:"disableNullMatches,omitempty"`
	Severities []string `json:"severities,omitempty"`
}
