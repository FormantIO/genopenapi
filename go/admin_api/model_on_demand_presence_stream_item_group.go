/*
 * Formant Cloud API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// OnDemandPresenceStreamItemGroup struct for OnDemandPresenceStreamItemGroup
type OnDemandPresenceStreamItemGroup struct {
	DatapointType string `json:"datapointType"`
	TimeRanges []OnDemandPresenceTimeRange `json:"timeRanges"`
}
