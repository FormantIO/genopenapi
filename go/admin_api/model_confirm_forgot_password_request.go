/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// ConfirmForgotPasswordRequest struct for ConfirmForgotPasswordRequest
type ConfirmForgotPasswordRequest struct {
	Email string `json:"email"`
	ConfirmationCode string `json:"confirmationCode"`
	NewPassword string `json:"newPassword"`
}
