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
// Organization struct for Organization
type Organization struct {
	Plan string `json:"plan"`
	Name string `json:"name"`
	Industry string `json:"industry"`
	Website string `json:"website"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City string `json:"city"`
	State string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country string `json:"country"`
	PagerdutyInfo PagerdutyInfo `json:"pagerdutyInfo,omitempty"`
	SlackInfo SlackInfo `json:"slackInfo,omitempty"`
	GoogleInfo GoogleInfo `json:"googleInfo,omitempty"`
	WebhooksInfo WebhooksInfo `json:"webhooksInfo,omitempty"`
	AwsInfo AwsInfo `json:"awsInfo,omitempty"`
	LookerInfo LookerInfo `json:"lookerInfo,omitempty"`
	StripeInfo StripeInfo `json:"stripeInfo,omitempty"`
	TeleopConfiguration UserTeleopConfiguration `json:"teleopConfiguration,omitempty"`
	TeleopEnabled bool `json:"teleopEnabled,omitempty"`
	AnalyticsEnabled bool `json:"analyticsEnabled,omitempty"`
	StripeBillingEnabled bool `json:"stripeBillingEnabled,omitempty"`
	StripeSubscriptionEnabled bool `json:"stripeSubscriptionEnabled,omitempty"`
	S3ExportEnabled bool `json:"s3ExportEnabled,omitempty"`
	BlobDataEnabled bool `json:"blobDataEnabled,omitempty"`
	DemoModeEnabled bool `json:"demoModeEnabled,omitempty"`
	DaysDataRetained int64 `json:"daysDataRetained,omitempty"`
	TrialPeriodEnd *time.Time `json:"trialPeriodEnd"`
	ExternalId string `json:"externalId,omitempty"`
	Id string `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
