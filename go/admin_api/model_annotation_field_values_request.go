/*
 * Formant Cloud API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// AnnotationFieldValuesRequest struct for AnnotationFieldValuesRequest
type AnnotationFieldValuesRequest struct {
	AnnotationTemplateId string `json:"annotationTemplateId"`
	Key interface{} `json:"key"`
	Tags map[string]string `json:"tags"`
}
