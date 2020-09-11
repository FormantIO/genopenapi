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
// AnnotationTemplate struct for AnnotationTemplate
type AnnotationTemplate struct {
	OrganizationId string `json:"organizationId,omitempty"`
	Name string `json:"name"`
	Tags map[string]string `json:"tags"`
	Description string `json:"description,omitempty"`
	Fields []AnnotationField `json:"fields"`
	Enabled bool `json:"enabled,omitempty"`
	PublishToGoogleSpreadsheetUrl *string `json:"publishToGoogleSpreadsheetUrl,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
