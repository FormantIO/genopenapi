/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// EventCounts struct for EventCounts
type EventCounts struct {
	Total int64 `json:"total"`
	Info int64 `json:"info,omitempty"`
	Warn int64 `json:"warn,omitempty"`
	Error int64 `json:"error,omitempty"`
	Critical int64 `json:"critical,omitempty"`
	TriggeredEvents int64 `json:"triggered-events,omitempty"`
	InterventionRequest int64 `json:"intervention-request,omitempty"`
}
