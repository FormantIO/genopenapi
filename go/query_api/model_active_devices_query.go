/*
 * Formant Cloud API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package query_api
import (
	"time"
)
// ActiveDevicesQuery struct for ActiveDevicesQuery
type ActiveDevicesQuery struct {
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
	OrganizationId string `json:"organizationId,omitempty"`
}
