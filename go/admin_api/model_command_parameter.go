/*
 * Formant admin-api
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
// CommandParameter struct for CommandParameter
type CommandParameter struct {
	Value string `json:"value,omitempty"`
	Meta map[string]string `json:"meta,omitempty"`
	ScrubberTime time.Time `json:"scrubberTime"`
}
