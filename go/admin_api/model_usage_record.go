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
// UsageRecord struct for UsageRecord
type UsageRecord struct {
	OrganizationId string `json:"organizationId"`
	Type string `json:"type"`
	Time time.Time `json:"time"`
	Value int64 `json:"value"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}