/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// LoginRequest struct for LoginRequest
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	TokenExpirationSeconds int64 `json:"tokenExpirationSeconds,omitempty"`
}
