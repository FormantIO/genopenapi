/*
 * Formant Cloud API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// BitsetAlertCondition struct for BitsetAlertCondition
type BitsetAlertCondition struct {
	Type string `json:"type"`
	BitConditions []BitCondition `json:"bitConditions"`
	Operator string `json:"operator"`
}
