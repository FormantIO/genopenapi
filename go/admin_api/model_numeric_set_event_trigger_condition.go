/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// NumericSetEventTriggerCondition struct for NumericSetEventTriggerCondition
type NumericSetEventTriggerCondition struct {
	Type string `json:"type"`
	Conditions []NumericCondition `json:"conditions"`
	Operator string `json:"operator"`
}