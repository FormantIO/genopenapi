/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// Filter struct for Filter
type Filter struct {
	AgentIds []string `json:"agentIds,omitempty"`
	DeviceIds []string `json:"deviceIds,omitempty"`
	Names []string `json:"names,omitempty"`
	Types []string `json:"types,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	NotNames []string `json:"notNames,omitempty"`
}
