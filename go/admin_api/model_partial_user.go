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
// PartialUser struct for PartialUser
type PartialUser struct {
	OrganizationId string `json:"organizationId,omitempty"`
	Email string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	Role string `json:"role,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	IsOrganizationOwner bool `json:"isOrganizationOwner,omitempty"`
	TermsAccepted string `json:"termsAccepted,omitempty"`
	LastLoggedIn time.Time `json:"lastLoggedIn,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
