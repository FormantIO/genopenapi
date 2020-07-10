/*
 * Formant Cloud API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech
 */

package query_api

	

// Query struct for Query
type Query struct {
	Start string `json:"start"`
	End string `json:"end"`
	Debug bool `json:"debug,omitempty"`
	AgentIds []string `json:"agentIds,omitempty"`
	DeviceIds []string `json:"deviceIds,omitempty"`
	Names []string `json:"names,omitempty"`
	Types []string `json:"types,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	NotNames []string `json:"notNames,omitempty"`
}
