/*
 * Formant admin-api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package admin_api
// HwInfo struct for HwInfo
type HwInfo struct {
	KernelInfo KernelInfo `json:"kernelInfo"`
	OsInfo OsInfo `json:"osInfo"`
	NodeInfo NodeInfo `json:"nodeInfo"`
	NetworkInfo NetworkInfo `json:"networkInfo"`
	HwEncodingAvailable bool `json:"hwEncodingAvailable"`
	VideoDevices []VideoDevice `json:"videoDevices,omitempty"`
	AudioCaptureDevices []AudioDevice `json:"audioCaptureDevices,omitempty"`
	OnvifDevices []OnvifDevice `json:"onvifDevices,omitempty"`
}
