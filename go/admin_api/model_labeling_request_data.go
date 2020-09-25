/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// LabelingRequestData struct for LabelingRequestData
type LabelingRequestData struct {
	Instruction string `json:"instruction"`
	ImageUrl string `json:"imageUrl"`
	Labels []Label `json:"labels"`
	Hint []LabeledPolygon `json:"hint,omitempty"`
	Title string `json:"title,omitempty"`
}
