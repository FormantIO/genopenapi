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
// DeviceLayout struct for DeviceLayout
type DeviceLayout struct {
	OrganizationId string `json:"organizationId,omitempty"`
	Name string `json:"name"`
	Tags map[string]string `json:"tags"`
	Filter Filter `json:"filter"`
	Layout interface{} `json:"layout"`
	Configuration []ViewConfiguration `json:"configuration"`
	Index int64 `json:"index"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
