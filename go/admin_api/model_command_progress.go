/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// CommandProgress struct for CommandProgress
type CommandProgress struct {
	CommandId string `json:"commandId"`
	Progress int64 `json:"progress"`
	Pending bool `json:"pending"`
}
