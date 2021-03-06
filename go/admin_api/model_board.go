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
// Board struct for Board
type Board struct {
	OrganizationId string `json:"organizationId,omitempty"`
	Name string `json:"name"`
	LookIds []float32 `json:"lookIds"`
	Layout interface{} `json:"layout"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
