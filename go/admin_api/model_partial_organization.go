/*
 * Formant Cloud API
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
// PartialOrganization struct for PartialOrganization
type PartialOrganization struct {
	Plan string `json:"plan,omitempty"`
	Name string `json:"name,omitempty"`
	Industry string `json:"industry,omitempty"`
	Website string `json:"website,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	Country string `json:"country,omitempty"`
	PagerdutyInfo PagerdutyInfo `json:"pagerdutyInfo,omitempty"`
	SlackInfo SlackInfo `json:"slackInfo,omitempty"`
	GoogleInfo GoogleInfo `json:"googleInfo,omitempty"`
	WebhooksInfo WebhooksInfo `json:"webhooksInfo,omitempty"`
	AwsInfo AwsInfo `json:"awsInfo,omitempty"`
	LookerInfo LookerInfo `json:"lookerInfo,omitempty"`
	StripeInfo StripeInfo `json:"stripeInfo,omitempty"`
	TeleopConfiguration UserTeleopConfiguration `json:"teleopConfiguration,omitempty"`
	TeleopEnabled bool `json:"teleopEnabled,omitempty"`
	StripeBillingEnabled bool `json:"stripeBillingEnabled,omitempty"`
	ExternalId string `json:"externalId,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
