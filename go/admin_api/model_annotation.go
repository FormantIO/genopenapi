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
// Annotation struct for Annotation
type Annotation struct {
	Type string `json:"type,omitempty"`
	EditedAt time.Time `json:"editedAt,omitempty"`
	UserId string `json:"userId"`
	AnnotationTemplateId string `json:"annotationTemplateId,omitempty"`
	TaggedUsers *interface{} `json:"taggedUsers,omitempty"`
	PublishedTo interface{} `json:"publishedTo,omitempty"`
	Note *string `json:"note,omitempty"`
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
