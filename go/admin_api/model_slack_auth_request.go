/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// SlackAuthRequest struct for SlackAuthRequest
type SlackAuthRequest struct {
	Code string `json:"code"`
	RedirectUri string `json:"redirectUri"`
}
