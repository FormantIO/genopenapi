/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// RefreshRequest struct for RefreshRequest
type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
	TokenExpirationSeconds int64 `json:"tokenExpirationSeconds,omitempty"`
}
