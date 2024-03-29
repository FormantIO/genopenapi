# Go API client for admin_api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

-
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

import the following package:

```golang
import "github.com/FormantIO/genopenapi/go/admin_api"
```

## Documentation for API Endpoints

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AnnotationApi* | [**AnnotationControllerDelete**](docs/AnnotationApi.md#annotationcontrollerdelete) | **Delete** /annotations/{id} | Delete
*AnnotationApi* | [**AnnotationControllerGetOne**](docs/AnnotationApi.md#annotationcontrollergetone) | **Get** /annotations/{id} | Get one
*AnnotationApi* | [**AnnotationControllerPatch**](docs/AnnotationApi.md#annotationcontrollerpatch) | **Patch** /annotations/{id} | Patch
*AnnotationApi* | [**AnnotationControllerPost**](docs/AnnotationApi.md#annotationcontrollerpost) | **Post** /annotations | Post
*AnnotationTemplateApi* | [**AnnotationTemplateControllerDelete**](docs/AnnotationTemplateApi.md#annotationtemplatecontrollerdelete) | **Delete** /annotation-templates/{id} | Delete
*AnnotationTemplateApi* | [**AnnotationTemplateControllerGetOne**](docs/AnnotationTemplateApi.md#annotationtemplatecontrollergetone) | **Get** /annotation-templates/{id} | Get one
*AnnotationTemplateApi* | [**AnnotationTemplateControllerInspectSpreadsheet**](docs/AnnotationTemplateApi.md#annotationtemplatecontrollerinspectspreadsheet) | **Post** /annotation-templates/inspect-spreadsheet | Inspect spreadsheet
*AnnotationTemplateApi* | [**AnnotationTemplateControllerList**](docs/AnnotationTemplateApi.md#annotationtemplatecontrollerlist) | **Get** /annotation-templates/ | List
*AnnotationTemplateApi* | [**AnnotationTemplateControllerPatch**](docs/AnnotationTemplateApi.md#annotationtemplatecontrollerpatch) | **Patch** /annotation-templates/{id} | Patch
*AnnotationTemplateApi* | [**AnnotationTemplateControllerPost**](docs/AnnotationTemplateApi.md#annotationtemplatecontrollerpost) | **Post** /annotation-templates | Post
*AnnotationTemplateApi* | [**AnnotationTemplateControllerQueryAnnotationFieldValues**](docs/AnnotationTemplateApi.md#annotationtemplatecontrollerqueryannotationfieldvalues) | **Post** /annotation-templates/query | Query annotation field values
*AuthApi* | [**AuthControllerChangePassword**](docs/AuthApi.md#authcontrollerchangepassword) | **Post** /auth/change-password | Change password
*AuthApi* | [**AuthControllerConfirmForgotPassword**](docs/AuthApi.md#authcontrollerconfirmforgotpassword) | **Post** /auth/confirm-forgot-password | Confirm forgot password
*AuthApi* | [**AuthControllerCreateServiceAccount**](docs/AuthApi.md#authcontrollercreateserviceaccount) | **Post** /auth/service-account | Create service account
*AuthApi* | [**AuthControllerDeviceCredentials**](docs/AuthApi.md#authcontrollerdevicecredentials) | **Post** /auth/device-credentials | Device credentials
*AuthApi* | [**AuthControllerForgotPassword**](docs/AuthApi.md#authcontrollerforgotpassword) | **Post** /auth/forgot-password | Forgot password
*AuthApi* | [**AuthControllerGetFeatures**](docs/AuthApi.md#authcontrollergetfeatures) | **Get** /auth/features | Get features
*AuthApi* | [**AuthControllerLogin**](docs/AuthApi.md#authcontrollerlogin) | **Post** /auth/login | Login
*AuthApi* | [**AuthControllerLoginGoogle**](docs/AuthApi.md#authcontrollerlogingoogle) | **Post** /auth/login-google | Login google
*AuthApi* | [**AuthControllerPlanUpgrade**](docs/AuthApi.md#authcontrollerplanupgrade) | **Post** /auth/upgrade-plan | Plan upgrade
*AuthApi* | [**AuthControllerRefresh**](docs/AuthApi.md#authcontrollerrefresh) | **Post** /auth/refresh | Refresh
*AuthApi* | [**AuthControllerResendConfirmationCode**](docs/AuthApi.md#authcontrollerresendconfirmationcode) | **Post** /auth/resend-confirmation-code | Resend confirmation code
*AuthApi* | [**AuthControllerResendInvitation**](docs/AuthApi.md#authcontrollerresendinvitation) | **Post** /auth/resend-invitation | Resend invitation
*AuthApi* | [**AuthControllerRespondToNewPasswordRequiredChallenge**](docs/AuthApi.md#authcontrollerrespondtonewpasswordrequiredchallenge) | **Post** /auth/respond-to-new-password-required-challenge | Respond to new password required challenge
*CaptureSessionApi* | [**CaptureSessionControllerAuthenticate**](docs/CaptureSessionApi.md#capturesessioncontrollerauthenticate) | **Post** /capture-sessions/{code}/authenticate | Authenticate
*CaptureSessionApi* | [**CaptureSessionControllerGetOne**](docs/CaptureSessionApi.md#capturesessioncontrollergetone) | **Get** /capture-sessions/{code} | Get one
*CaptureSessionApi* | [**CaptureSessionControllerPost**](docs/CaptureSessionApi.md#capturesessioncontrollerpost) | **Post** /capture-sessions | Post
*ChannelApi* | [**ChannelControllerDelete**](docs/ChannelApi.md#channelcontrollerdelete) | **Delete** /channels/{id} | Delete
*ChannelApi* | [**ChannelControllerGetAll**](docs/ChannelApi.md#channelcontrollergetall) | **Get** /channels | Get all
*ChannelApi* | [**ChannelControllerGetOne**](docs/ChannelApi.md#channelcontrollergetone) | **Get** /channels/{id} | Get one
*ChannelApi* | [**ChannelControllerPatch**](docs/ChannelApi.md#channelcontrollerpatch) | **Patch** /channels/{id} | Patch
*ChannelApi* | [**ChannelControllerPost**](docs/ChannelApi.md#channelcontrollerpost) | **Post** /channels | Post
*CommandApi* | [**CommandControllerGetOne**](docs/CommandApi.md#commandcontrollergetone) | **Get** /commands/{id} | Get one
*CommandApi* | [**CommandControllerPatch**](docs/CommandApi.md#commandcontrollerpatch) | **Patch** /commands/{id} | Patch
*CommandApi* | [**CommandControllerPost**](docs/CommandApi.md#commandcontrollerpost) | **Post** /commands | Post
*CommandApi* | [**CommandControllerQuery**](docs/CommandApi.md#commandcontrollerquery) | **Post** /commands/query | Query
*CommandRequestApi* | [**CommandRequestControllerPoll**](docs/CommandRequestApi.md#commandrequestcontrollerpoll) | **Post** /command-requests/pending | Poll
*CommandResponseApi* | [**CommandResponseControllerPost**](docs/CommandResponseApi.md#commandresponsecontrollerpost) | **Post** /command-responses | Post
*CommandTemplateApi* | [**CommandTemplateControllerDelete**](docs/CommandTemplateApi.md#commandtemplatecontrollerdelete) | **Delete** /command-templates/{id} | Delete
*CommandTemplateApi* | [**CommandTemplateControllerGetOne**](docs/CommandTemplateApi.md#commandtemplatecontrollergetone) | **Get** /command-templates/{id} | Get one
*CommandTemplateApi* | [**CommandTemplateControllerList**](docs/CommandTemplateApi.md#commandtemplatecontrollerlist) | **Get** /command-templates/ | List
*CommandTemplateApi* | [**CommandTemplateControllerPatch**](docs/CommandTemplateApi.md#commandtemplatecontrollerpatch) | **Patch** /command-templates/{id} | Patch
*CommandTemplateApi* | [**CommandTemplateControllerPost**](docs/CommandTemplateApi.md#commandtemplatecontrollerpost) | **Post** /command-templates | Post
*CommentApi* | [**CommentControllerDelete**](docs/CommentApi.md#commentcontrollerdelete) | **Delete** /comments/{id} | Delete
*CommentApi* | [**CommentControllerGetOne**](docs/CommentApi.md#commentcontrollergetone) | **Get** /comments/{id} | Get one
*CommentApi* | [**CommentControllerPatch**](docs/CommentApi.md#commentcontrollerpatch) | **Patch** /comments/{id} | Patch
*CommentApi* | [**CommentControllerPost**](docs/CommentApi.md#commentcontrollerpost) | **Post** /comments | Post
*CustomEventApi* | [**CustomEventControllerPost**](docs/CustomEventApi.md#customeventcontrollerpost) | **Post** /custom-events | Post
*DeviceApi* | [**DeviceControllerCount**](docs/DeviceApi.md#devicecontrollercount) | **Post** /devices/count | Count
*DeviceApi* | [**DeviceControllerDisable**](docs/DeviceApi.md#devicecontrollerdisable) | **Post** /devices/{id}/disable | Disable
*DeviceApi* | [**DeviceControllerFilter**](docs/DeviceApi.md#devicecontrollerfilter) | **Post** /devices/filter | Filter
*DeviceApi* | [**DeviceControllerGenerateDeviceProvisioningToken**](docs/DeviceApi.md#devicecontrollergeneratedeviceprovisioningtoken) | **Post** /devices/{id}/provisioning-token | Generate device provisioning token
*DeviceApi* | [**DeviceControllerGetConfiguration**](docs/DeviceApi.md#devicecontrollergetconfiguration) | **Get** /devices/{id}/configurations/{version} | Get configuration
*DeviceApi* | [**DeviceControllerGetGeoIp**](docs/DeviceApi.md#devicecontrollergetgeoip) | **Get** /devices/{id}/geoip | Get geo ip
*DeviceApi* | [**DeviceControllerGetOne**](docs/DeviceApi.md#devicecontrollergetone) | **Get** /devices/{id} | Get one
*DeviceApi* | [**DeviceControllerGetTags**](docs/DeviceApi.md#devicecontrollergettags) | **Post** /devices/tags | Get tags
*DeviceApi* | [**DeviceControllerGetUpdatedAgentVersion**](docs/DeviceApi.md#devicecontrollergetupdatedagentversion) | **Get** /devices/{id}/updated-agent-version | Get updated agent version
*DeviceApi* | [**DeviceControllerGetUpdatedConfiguration**](docs/DeviceApi.md#devicecontrollergetupdatedconfiguration) | **Get** /devices/{id}/updated-configuration | Get updated configuration
*DeviceApi* | [**DeviceControllerPatch**](docs/DeviceApi.md#devicecontrollerpatch) | **Patch** /devices/{id} | Patch
*DeviceApi* | [**DeviceControllerPost**](docs/DeviceApi.md#devicecontrollerpost) | **Post** /devices | Post
*DeviceApi* | [**DeviceControllerPostConfiguration**](docs/DeviceApi.md#devicecontrollerpostconfiguration) | **Post** /devices/{id}/configurations | Post configuration
*DeviceApi* | [**DeviceControllerProvisionDevice**](docs/DeviceApi.md#devicecontrollerprovisiondevice) | **Post** /devices/provision | Provision device
*DeviceApi* | [**DeviceControllerQuery**](docs/DeviceApi.md#devicecontrollerquery) | **Post** /devices/query | Query
*DeviceApi* | [**DeviceControllerUnprovisionDevice**](docs/DeviceApi.md#devicecontrollerunprovisiondevice) | **Post** /devices/{id}/unprovision | Unprovision device
*DeviceApi* | [**DeviceControllerValidateStreamConfiguration**](docs/DeviceApi.md#devicecontrollervalidatestreamconfiguration) | **Post** /devices/validate-stream-configuration | Validate stream configuration
*DeviceConfigurationTemplateApi* | [**DeviceConfigurationTemplateControllerDelete**](docs/DeviceConfigurationTemplateApi.md#deviceconfigurationtemplatecontrollerdelete) | **Delete** /device-configuration-templates/{id} | Delete
*DeviceConfigurationTemplateApi* | [**DeviceConfigurationTemplateControllerGetOne**](docs/DeviceConfigurationTemplateApi.md#deviceconfigurationtemplatecontrollergetone) | **Get** /device-configuration-templates/{id} | Get one
*DeviceConfigurationTemplateApi* | [**DeviceConfigurationTemplateControllerList**](docs/DeviceConfigurationTemplateApi.md#deviceconfigurationtemplatecontrollerlist) | **Get** /device-configuration-templates/ | List
*DeviceConfigurationTemplateApi* | [**DeviceConfigurationTemplateControllerPatch**](docs/DeviceConfigurationTemplateApi.md#deviceconfigurationtemplatecontrollerpatch) | **Patch** /device-configuration-templates/{id} | Patch
*DeviceConfigurationTemplateApi* | [**DeviceConfigurationTemplateControllerPost**](docs/DeviceConfigurationTemplateApi.md#deviceconfigurationtemplatecontrollerpost) | **Post** /device-configuration-templates | Post
*DeviceDetailsApi* | [**DeviceDetailsControllerGetAll**](docs/DeviceDetailsApi.md#devicedetailscontrollergetall) | **Post** /device-details/query | Get all
*DeviceDetailsApi* | [**DeviceDetailsControllerGetOne**](docs/DeviceDetailsApi.md#devicedetailscontrollergetone) | **Get** /device-details/{id} | Get one
*EventApi* | [**EventControllerAnnotationTemplates**](docs/EventApi.md#eventcontrollerannotationtemplates) | **Post** /events/annotation-templates | Annotation templates
*EventApi* | [**EventControllerCount**](docs/EventApi.md#eventcontrollercount) | **Post** /events/count | Count
*EventApi* | [**EventControllerCounts**](docs/EventApi.md#eventcontrollercounts) | **Post** /events/counts | Counts
*EventApi* | [**EventControllerDevices**](docs/EventApi.md#eventcontrollerdevices) | **Post** /events/devices | Devices
*EventApi* | [**EventControllerEventTypes**](docs/EventApi.md#eventcontrollereventtypes) | **Post** /events/event-types | Event types
*EventApi* | [**EventControllerExportSheet**](docs/EventApi.md#eventcontrollerexportsheet) | **Post** /events/export-sheet | Export sheet
*EventApi* | [**EventControllerGetOne**](docs/EventApi.md#eventcontrollergetone) | **Get** /events/{id} | Get one
*EventApi* | [**EventControllerHistogram**](docs/EventApi.md#eventcontrollerhistogram) | **Post** /events/histogram | Histogram
*EventApi* | [**EventControllerQuery**](docs/EventApi.md#eventcontrollerquery) | **Post** /events/query | Query
*EventApi* | [**EventControllerSeek**](docs/EventApi.md#eventcontrollerseek) | **Post** /events/seek | Seek
*EventApi* | [**EventControllerSeverities**](docs/EventApi.md#eventcontrollerseverities) | **Post** /events/severities | Severities
*EventApi* | [**EventControllerStreamNames**](docs/EventApi.md#eventcontrollerstreamnames) | **Post** /events/stream-names | Stream names
*EventApi* | [**EventControllerStreamTypes**](docs/EventApi.md#eventcontrollerstreamtypes) | **Post** /events/stream-types | Stream types
*EventApi* | [**EventControllerTags**](docs/EventApi.md#eventcontrollertags) | **Post** /events/tags | Tags
*EventApi* | [**EventControllerUsers**](docs/EventApi.md#eventcontrollerusers) | **Post** /events/users | Users
*EventApi* | [**EventControllerView**](docs/EventApi.md#eventcontrollerview) | **Post** /events/view | View
*EventTriggerApi* | [**EventTriggerControllerGetAll**](docs/EventTriggerApi.md#eventtriggercontrollergetall) | **Get** /event-triggers | Get all
*EventTriggerApi* | [**EventTriggerControllerGetOne**](docs/EventTriggerApi.md#eventtriggercontrollergetone) | **Get** /event-triggers/{id} | Get one
*EventTriggerApi* | [**EventTriggerControllerPatch**](docs/EventTriggerApi.md#eventtriggercontrollerpatch) | **Patch** /event-triggers/{id} | Patch
*EventTriggerApi* | [**EventTriggerControllerPost**](docs/EventTriggerApi.md#eventtriggercontrollerpost) | **Post** /event-triggers | Post
*EventTriggerApi* | [**EventTriggerControllerQuery**](docs/EventTriggerApi.md#eventtriggercontrollerquery) | **Post** /event-triggers/updated | Query
*FileApi* | [**FileControllerBeginUpload**](docs/FileApi.md#filecontrollerbeginupload) | **Post** /files/begin-upload | Begin upload
*FileApi* | [**FileControllerCompleteUpload**](docs/FileApi.md#filecontrollercompleteupload) | **Post** /files/complete-upload | Complete upload
*FileApi* | [**FileControllerQuery**](docs/FileApi.md#filecontrollerquery) | **Post** /files/query | Query
*IntegrationApi* | [**IntegrationControllerGoogle**](docs/IntegrationApi.md#integrationcontrollergoogle) | **Post** /integrations/google/auth | Google
*IntegrationApi* | [**IntegrationControllerSlack**](docs/IntegrationApi.md#integrationcontrollerslack) | **Post** /integrations/slack/auth | Slack
*InterventionRequestApi* | [**InterventionRequestControllerGetOne**](docs/InterventionRequestApi.md#interventionrequestcontrollergetone) | **Get** /intervention-requests/{id} | Get one
*InterventionRequestApi* | [**InterventionRequestControllerPost**](docs/InterventionRequestApi.md#interventionrequestcontrollerpost) | **Post** /intervention-requests | Post
*InterventionResponseApi* | [**InterventionResponseControllerPost**](docs/InterventionResponseApi.md#interventionresponsecontrollerpost) | **Post** /intervention-responses | Post
*OrganizationApi* | [**OrganizationControllerQueryUsageRecords**](docs/OrganizationApi.md#organizationcontrollerqueryusagerecords) | **Post** /organizations/{id}/usage-record/query | Query usage records
*ShareApi* | [**ShareControllerAuthenticate**](docs/ShareApi.md#sharecontrollerauthenticate) | **Post** /shares/{code}/authenticate | Authenticate
*ShareApi* | [**ShareControllerDelete**](docs/ShareApi.md#sharecontrollerdelete) | **Delete** /shares/{code} | Delete
*ShareApi* | [**ShareControllerGetAll**](docs/ShareApi.md#sharecontrollergetall) | **Get** /shares | Get all
*ShareApi* | [**ShareControllerGetOne**](docs/ShareApi.md#sharecontrollergetone) | **Get** /shares/{code} | Get one
*ShareApi* | [**ShareControllerPost**](docs/ShareApi.md#sharecontrollerpost) | **Post** /shares | Post
*TriggeredEventApi* | [**TriggeredEventControllerGetOne**](docs/TriggeredEventApi.md#triggeredeventcontrollergetone) | **Get** /triggered-events/{id} | Get one
*UserApi* | [**UserControllerGetAll**](docs/UserApi.md#usercontrollergetall) | **Get** /users | Get all
*UserApi* | [**UserControllerGetOne**](docs/UserApi.md#usercontrollergetone) | **Get** /users/{id} | Get one
*UserApi* | [**UserControllerPatch**](docs/UserApi.md#usercontrollerpatch) | **Patch** /users/{id} | Patch
*UserApi* | [**UserControllerPost**](docs/UserApi.md#usercontrollerpost) | **Post** /users | Post
*ViewApi* | [**ViewControllerDelete**](docs/ViewApi.md#viewcontrollerdelete) | **Delete** /views/{id} | Delete
*ViewApi* | [**ViewControllerGetAll**](docs/ViewApi.md#viewcontrollergetall) | **Get** /views | Get all
*ViewApi* | [**ViewControllerGetOne**](docs/ViewApi.md#viewcontrollergetone) | **Get** /views/{id} | Get one
*ViewApi* | [**ViewControllerPatch**](docs/ViewApi.md#viewcontrollerpatch) | **Patch** /views/{id} | Patch
*ViewApi* | [**ViewControllerPost**](docs/ViewApi.md#viewcontrollerpost) | **Post** /views | Post


## Documentation For Models

 - [Adapter](docs/Adapter.md)
 - [AdapterConfiguration](docs/AdapterConfiguration.md)
 - [Annotation](docs/Annotation.md)
 - [AnnotationField](docs/AnnotationField.md)
 - [AnnotationFieldValue](docs/AnnotationFieldValue.md)
 - [AnnotationFieldValuesRequest](docs/AnnotationFieldValuesRequest.md)
 - [AnnotationFieldValuesResponse](docs/AnnotationFieldValuesResponse.md)
 - [AnnotationTemplate](docs/AnnotationTemplate.md)
 - [AnnotationTemplateListResponse](docs/AnnotationTemplateListResponse.md)
 - [AudioDevice](docs/AudioDevice.md)
 - [Authentication](docs/Authentication.md)
 - [AwsInfo](docs/AwsInfo.md)
 - [BaseEvent](docs/BaseEvent.md)
 - [BatteryEventTriggerCondition](docs/BatteryEventTriggerCondition.md)
 - [BeginUploadRequest](docs/BeginUploadRequest.md)
 - [BeginUploadResponse](docs/BeginUploadResponse.md)
 - [BillingInfo](docs/BillingInfo.md)
 - [BitCondition](docs/BitCondition.md)
 - [BitsetEventTriggerCondition](docs/BitsetEventTriggerCondition.md)
 - [BitsetViewConfiguration](docs/BitsetViewConfiguration.md)
 - [Board](docs/Board.md)
 - [Camera](docs/Camera.md)
 - [CaptureSession](docs/CaptureSession.md)
 - [Challenge](docs/Challenge.md)
 - [ChangePasswordRequest](docs/ChangePasswordRequest.md)
 - [Channel](docs/Channel.md)
 - [ChannelListResponse](docs/ChannelListResponse.md)
 - [Command](docs/Command.md)
 - [CommandListResponse](docs/CommandListResponse.md)
 - [CommandParameter](docs/CommandParameter.md)
 - [CommandProgress](docs/CommandProgress.md)
 - [CommandQuery](docs/CommandQuery.md)
 - [CommandRequest](docs/CommandRequest.md)
 - [CommandRequestListResponse](docs/CommandRequestListResponse.md)
 - [CommandResponse](docs/CommandResponse.md)
 - [CommandTemplate](docs/CommandTemplate.md)
 - [CommandTemplateListResponse](docs/CommandTemplateListResponse.md)
 - [Comment](docs/Comment.md)
 - [CompleteUploadRequest](docs/CompleteUploadRequest.md)
 - [ConfirmForgotPasswordRequest](docs/ConfirmForgotPasswordRequest.md)
 - [CreateServiceAccountRequest](docs/CreateServiceAccountRequest.md)
 - [CreateServiceAccountResponse](docs/CreateServiceAccountResponse.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CustomEvent](docs/CustomEvent.md)
 - [Device](docs/Device.md)
 - [DeviceApplicationConfiguration](docs/DeviceApplicationConfiguration.md)
 - [DeviceBlobData](docs/DeviceBlobData.md)
 - [DeviceConfiguration](docs/DeviceConfiguration.md)
 - [DeviceConfigurationDocument](docs/DeviceConfigurationDocument.md)
 - [DeviceConfigurationTemplate](docs/DeviceConfigurationTemplate.md)
 - [DeviceConfigurationTemplateListResponse](docs/DeviceConfigurationTemplateListResponse.md)
 - [DeviceCredentials](docs/DeviceCredentials.md)
 - [DeviceDetails](docs/DeviceDetails.md)
 - [DeviceDetailsListResponse](docs/DeviceDetailsListResponse.md)
 - [DeviceDiagnosticsConfiguration](docs/DeviceDiagnosticsConfiguration.md)
 - [DeviceDiskConfiguration](docs/DeviceDiskConfiguration.md)
 - [DeviceFollower](docs/DeviceFollower.md)
 - [DeviceListResponse](docs/DeviceListResponse.md)
 - [DevicePortForwardingConfiguration](docs/DevicePortForwardingConfiguration.md)
 - [DeviceProvisioning](docs/DeviceProvisioning.md)
 - [DeviceProvisioningRequest](docs/DeviceProvisioningRequest.md)
 - [DeviceQuery](docs/DeviceQuery.md)
 - [DeviceReportedConfigurationState](docs/DeviceReportedConfigurationState.md)
 - [DeviceResourcesConfiguration](docs/DeviceResourcesConfiguration.md)
 - [DeviceRosConfiguration](docs/DeviceRosConfiguration.md)
 - [DeviceRosState](docs/DeviceRosState.md)
 - [DeviceState](docs/DeviceState.md)
 - [DeviceStreamConfiguration](docs/DeviceStreamConfiguration.md)
 - [DeviceStreamCustomConfiguration](docs/DeviceStreamCustomConfiguration.md)
 - [DeviceStreamDirectoryWatchConfiguration](docs/DeviceStreamDirectoryWatchConfiguration.md)
 - [DeviceStreamFileTailConfiguration](docs/DeviceStreamFileTailConfiguration.md)
 - [DeviceStreamHardwareConfiguration](docs/DeviceStreamHardwareConfiguration.md)
 - [DeviceStreamRosLocalizationConfiguration](docs/DeviceStreamRosLocalizationConfiguration.md)
 - [DeviceStreamRosTopicConfiguration](docs/DeviceStreamRosTopicConfiguration.md)
 - [DeviceStreamRosTransformTreeConfiguration](docs/DeviceStreamRosTransformTreeConfiguration.md)
 - [DeviceStreamTransformConfiguration](docs/DeviceStreamTransformConfiguration.md)
 - [DeviceTelemetryConfiguration](docs/DeviceTelemetryConfiguration.md)
 - [DeviceTeleopConfiguration](docs/DeviceTeleopConfiguration.md)
 - [DeviceTeleopCustomStreamConfiguration](docs/DeviceTeleopCustomStreamConfiguration.md)
 - [DeviceTeleopHardwareStreamConfiguration](docs/DeviceTeleopHardwareStreamConfiguration.md)
 - [DeviceTeleopRosStreamConfiguration](docs/DeviceTeleopRosStreamConfiguration.md)
 - [EventCounts](docs/EventCounts.md)
 - [EventExportSheetRequest](docs/EventExportSheetRequest.md)
 - [EventExportSheetResult](docs/EventExportSheetResult.md)
 - [EventFilter](docs/EventFilter.md)
 - [EventHistogram](docs/EventHistogram.md)
 - [EventHistogramEntry](docs/EventHistogramEntry.md)
 - [EventListResponse](docs/EventListResponse.md)
 - [EventQuery](docs/EventQuery.md)
 - [EventSeekQuery](docs/EventSeekQuery.md)
 - [EventSort](docs/EventSort.md)
 - [EventTrigger](docs/EventTrigger.md)
 - [EventTriggerCommand](docs/EventTriggerCommand.md)
 - [EventTriggerListResponse](docs/EventTriggerListResponse.md)
 - [FileInfo](docs/FileInfo.md)
 - [Filter](docs/Filter.md)
 - [ForgotPasswordRequest](docs/ForgotPasswordRequest.md)
 - [ForwardingConfiguration](docs/ForwardingConfiguration.md)
 - [GeoIp](docs/GeoIp.md)
 - [GetFeaturesResponse](docs/GetFeaturesResponse.md)
 - [GoogleAuthRequest](docs/GoogleAuthRequest.md)
 - [GoogleInfo](docs/GoogleInfo.md)
 - [GoogleLoginRequest](docs/GoogleLoginRequest.md)
 - [GoogleSheetParseResult](docs/GoogleSheetParseResult.md)
 - [GoogleSpreadsheetInspection](docs/GoogleSpreadsheetInspection.md)
 - [GoogleStorageExport](docs/GoogleStorageExport.md)
 - [GoogleStorageInfo](docs/GoogleStorageInfo.md)
 - [Group](docs/Group.md)
 - [HwInfo](docs/HwInfo.md)
 - [ImageViewConfiguration](docs/ImageViewConfiguration.md)
 - [InspectSpreadsheetRequest](docs/InspectSpreadsheetRequest.md)
 - [InspectSpreadsheetResponse](docs/InspectSpreadsheetResponse.md)
 - [IntervalEventFilter](docs/IntervalEventFilter.md)
 - [InterventionRequest](docs/InterventionRequest.md)
 - [InterventionResponse](docs/InterventionResponse.md)
 - [JoystickConfiguration](docs/JoystickConfiguration.md)
 - [KernelInfo](docs/KernelInfo.md)
 - [Label](docs/Label.md)
 - [LabeledPolygon](docs/LabeledPolygon.md)
 - [LabelingRequestData](docs/LabelingRequestData.md)
 - [LocalizationViewConfiguration](docs/LocalizationViewConfiguration.md)
 - [Location](docs/Location.md)
 - [LocationViewConfiguration](docs/LocationViewConfiguration.md)
 - [LocationViewport](docs/LocationViewport.md)
 - [LoginRequest](docs/LoginRequest.md)
 - [LoginResult](docs/LoginResult.md)
 - [LookerInfo](docs/LookerInfo.md)
 - [LookerLook](docs/LookerLook.md)
 - [Network](docs/Network.md)
 - [NetworkInfo](docs/NetworkInfo.md)
 - [NodeInfo](docs/NodeInfo.md)
 - [NumericCondition](docs/NumericCondition.md)
 - [NumericSetEventTriggerCondition](docs/NumericSetEventTriggerCondition.md)
 - [NumericViewConfiguration](docs/NumericViewConfiguration.md)
 - [OnDemandBuffer](docs/OnDemandBuffer.md)
 - [OnDemandPresenceStreamItemGroup](docs/OnDemandPresenceStreamItemGroup.md)
 - [OnDemandPresenceTimeRange](docs/OnDemandPresenceTimeRange.md)
 - [OnDemandState](docs/OnDemandState.md)
 - [OnDemandStreamPresence](docs/OnDemandStreamPresence.md)
 - [OnvifDevice](docs/OnvifDevice.md)
 - [Organization](docs/Organization.md)
 - [OsInfo](docs/OsInfo.md)
 - [PagerdutyInfo](docs/PagerdutyInfo.md)
 - [PartialAnnotation](docs/PartialAnnotation.md)
 - [PartialAnnotationTemplate](docs/PartialAnnotationTemplate.md)
 - [PartialChannel](docs/PartialChannel.md)
 - [PartialCommand](docs/PartialCommand.md)
 - [PartialCommandTemplate](docs/PartialCommandTemplate.md)
 - [PartialComment](docs/PartialComment.md)
 - [PartialDevice](docs/PartialDevice.md)
 - [PartialDeviceConfigurationTemplate](docs/PartialDeviceConfigurationTemplate.md)
 - [PartialEventTrigger](docs/PartialEventTrigger.md)
 - [PartialUser](docs/PartialUser.md)
 - [PartialView](docs/PartialView.md)
 - [PointCloudViewConfiguration](docs/PointCloudViewConfiguration.md)
 - [PollCommandRequest](docs/PollCommandRequest.md)
 - [PresenceEventTriggerCondition](docs/PresenceEventTriggerCondition.md)
 - [QueryFilesRequest](docs/QueryFilesRequest.md)
 - [QueryFilesResponse](docs/QueryFilesResponse.md)
 - [RefreshRequest](docs/RefreshRequest.md)
 - [RegexEventTriggerCondition](docs/RegexEventTriggerCondition.md)
 - [ResendConfirmationCodeRequest](docs/ResendConfirmationCodeRequest.md)
 - [ResendInvitationRequest](docs/ResendInvitationRequest.md)
 - [RespondToNewPasswordRequiredChallengeRequest](docs/RespondToNewPasswordRequiredChallengeRequest.md)
 - [RosTopic](docs/RosTopic.md)
 - [RtcInfo](docs/RtcInfo.md)
 - [S3Export](docs/S3Export.md)
 - [ScopeFilter](docs/ScopeFilter.md)
 - [SelectionRequestData](docs/SelectionRequestData.md)
 - [Share](docs/Share.md)
 - [ShareListResponse](docs/ShareListResponse.md)
 - [SheetParameters](docs/SheetParameters.md)
 - [SlackAuthRequest](docs/SlackAuthRequest.md)
 - [SlackInfo](docs/SlackInfo.md)
 - [Stream](docs/Stream.md)
 - [StringListResponse](docs/StringListResponse.md)
 - [StripeCard](docs/StripeCard.md)
 - [StripeInfo](docs/StripeInfo.md)
 - [TagParameters](docs/TagParameters.md)
 - [TagTemplate](docs/TagTemplate.md)
 - [TagsResponse](docs/TagsResponse.md)
 - [TeleopJoystickAxisConfiguration](docs/TeleopJoystickAxisConfiguration.md)
 - [TeleopJoystickConfiguration](docs/TeleopJoystickConfiguration.md)
 - [TeleopViewConfiguration](docs/TeleopViewConfiguration.md)
 - [ThresholdEventTriggerCondition](docs/ThresholdEventTriggerCondition.md)
 - [TokenResult](docs/TokenResult.md)
 - [TransformTreeViewConfiguration](docs/TransformTreeViewConfiguration.md)
 - [TriggeredConfiguration](docs/TriggeredConfiguration.md)
 - [TriggeredEvent](docs/TriggeredEvent.md)
 - [UpdatedAgentVersionResponse](docs/UpdatedAgentVersionResponse.md)
 - [UpdatedConfigurationResponse](docs/UpdatedConfigurationResponse.md)
 - [UpdatedEventTriggerRequest](docs/UpdatedEventTriggerRequest.md)
 - [UpdatedEventTriggerResponse](docs/UpdatedEventTriggerResponse.md)
 - [UsagePrices](docs/UsagePrices.md)
 - [UsageRecord](docs/UsageRecord.md)
 - [UsageRecordQuery](docs/UsageRecordQuery.md)
 - [UsageRecordQueryResponse](docs/UsageRecordQueryResponse.md)
 - [User](docs/User.md)
 - [UserListResponse](docs/UserListResponse.md)
 - [UserParameters](docs/UserParameters.md)
 - [UserTeleopConfiguration](docs/UserTeleopConfiguration.md)
 - [UserTeleopRosStreamConfiguration](docs/UserTeleopRosStreamConfiguration.md)
 - [UserTeleopTwistRosTopicConfiguration](docs/UserTeleopTwistRosTopicConfiguration.md)
 - [UuidListResponse](docs/UuidListResponse.md)
 - [Vector3](docs/Vector3.md)
 - [VideoDevice](docs/VideoDevice.md)
 - [View](docs/View.md)
 - [ViewConfiguration](docs/ViewConfiguration.md)
 - [Webhook](docs/Webhook.md)
 - [WebhooksInfo](docs/WebhooksInfo.md)


## Documentation For Authorization



## bearerAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```



## Author



