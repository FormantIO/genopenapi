/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// DeviceDetails struct for DeviceDetails
type DeviceDetails struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Tags map[string]string `json:"tags"`
	Enabled bool `json:"enabled"`
	PublicKey string `json:"publicKey"`
}
