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
// GoogleInfo struct for GoogleInfo
type GoogleInfo struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken string `json:"accessToken"`
	ExpirationDate time.Time `json:"expirationDate"`
	Scope string `json:"scope"`
}
