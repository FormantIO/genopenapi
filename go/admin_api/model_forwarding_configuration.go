/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// ForwardingConfiguration struct for ForwardingConfiguration
type ForwardingConfiguration struct {
	Pagerduty bool `json:"pagerduty,omitempty"`
	Slack bool `json:"slack,omitempty"`
	Webhooks []string `json:"webhooks,omitempty"`
	SmsFollowers []string `json:"smsFollowers,omitempty"`
}