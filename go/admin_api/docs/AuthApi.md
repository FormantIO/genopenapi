# \AuthApi

All URIs are relative to *https://api-dev.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthControllerChangePassword**](AuthApi.md#AuthControllerChangePassword) | **Post** /auth/change-password | Change password
[**AuthControllerConfirmForgotPassword**](AuthApi.md#AuthControllerConfirmForgotPassword) | **Post** /auth/confirm-forgot-password | Confirm forgot password
[**AuthControllerCreateServiceAccount**](AuthApi.md#AuthControllerCreateServiceAccount) | **Post** /auth/service-account | Create service account
[**AuthControllerDeviceCredentials**](AuthApi.md#AuthControllerDeviceCredentials) | **Post** /auth/device-credentials | Device credentials
[**AuthControllerForgotPassword**](AuthApi.md#AuthControllerForgotPassword) | **Post** /auth/forgot-password | Forgot password
[**AuthControllerGetFeatures**](AuthApi.md#AuthControllerGetFeatures) | **Get** /auth/features | Get features
[**AuthControllerLogin**](AuthApi.md#AuthControllerLogin) | **Post** /auth/login | Login
[**AuthControllerLoginGoogle**](AuthApi.md#AuthControllerLoginGoogle) | **Post** /auth/login-google | Login google
[**AuthControllerPlanUpgrade**](AuthApi.md#AuthControllerPlanUpgrade) | **Post** /auth/upgrade-plan | Plan upgrade
[**AuthControllerRefresh**](AuthApi.md#AuthControllerRefresh) | **Post** /auth/refresh | Refresh
[**AuthControllerResendConfirmationCode**](AuthApi.md#AuthControllerResendConfirmationCode) | **Post** /auth/resend-confirmation-code | Resend confirmation code
[**AuthControllerResendInvitation**](AuthApi.md#AuthControllerResendInvitation) | **Post** /auth/resend-invitation | Resend invitation
[**AuthControllerRespondToNewPasswordRequiredChallenge**](AuthApi.md#AuthControllerRespondToNewPasswordRequiredChallenge) | **Post** /auth/respond-to-new-password-required-challenge | Respond to new password required challenge



## AuthControllerChangePassword

> AuthControllerChangePassword(ctx, changePasswordRequest)

Change password

Change Password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**changePasswordRequest** | [**ChangePasswordRequest**](ChangePasswordRequest.md)| ChangePasswordRequest | 

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerConfirmForgotPassword

> AuthControllerConfirmForgotPassword(ctx, confirmForgotPasswordRequest)

Confirm forgot password

Confirm Forgot Password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**confirmForgotPasswordRequest** | [**ConfirmForgotPasswordRequest**](ConfirmForgotPasswordRequest.md)| ConfirmForgotPasswordRequest | 

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerCreateServiceAccount

> CreateServiceAccountResponse AuthControllerCreateServiceAccount(ctx, createServiceAccountRequest)

Create service account

Create a service account Authorized clients: administrator Authorized plans: starter, premium

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**createServiceAccountRequest** | [**CreateServiceAccountRequest**](CreateServiceAccountRequest.md)| CreateServiceAccountRequest | 

### Return type

[**CreateServiceAccountResponse**](CreateServiceAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerDeviceCredentials

> DeviceCredentials AuthControllerDeviceCredentials(ctx, )

Device credentials

Device Credentials Authorized clients: device Authorized plans: freemium, starter, premium

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DeviceCredentials**](DeviceCredentials.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerForgotPassword

> AuthControllerForgotPassword(ctx, forgotPasswordRequest)

Forgot password

Forgot Password

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**forgotPasswordRequest** | [**ForgotPasswordRequest**](ForgotPasswordRequest.md)| ForgotPasswordRequest | 

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerGetFeatures

> GetFeaturesResponse AuthControllerGetFeatures(ctx, )

Get features

Get enabled features Authorized clients: viewer, device Authorized plans: freemium, starter, premium

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetFeaturesResponse**](GetFeaturesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerLogin

> LoginResult AuthControllerLogin(ctx, loginRequest)

Login

Login

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loginRequest** | [**LoginRequest**](LoginRequest.md)| LoginRequest | 

### Return type

[**LoginResult**](LoginResult.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerLoginGoogle

> LoginResult AuthControllerLoginGoogle(ctx, googleLoginRequest)

Login google

Login with Google Token

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**googleLoginRequest** | [**GoogleLoginRequest**](GoogleLoginRequest.md)| GoogleLoginRequest | 

### Return type

[**LoginResult**](LoginResult.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerPlanUpgrade

> AuthControllerPlanUpgrade(ctx, )

Plan upgrade

Signal intent to upgrade Authorized clients: viewer Authorized plans: freemium, starter, premium

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerRefresh

> LoginResult AuthControllerRefresh(ctx, refreshRequest)

Refresh

Refresh

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refreshRequest** | [**RefreshRequest**](RefreshRequest.md)| RefreshRequest | 

### Return type

[**LoginResult**](LoginResult.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerResendConfirmationCode

> AuthControllerResendConfirmationCode(ctx, resendConfirmationCodeRequest)

Resend confirmation code

Resend Confirmation Code

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resendConfirmationCodeRequest** | [**ResendConfirmationCodeRequest**](ResendConfirmationCodeRequest.md)| ResendConfirmationCodeRequest | 

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerResendInvitation

> AuthControllerResendInvitation(ctx, resendInvitationRequest)

Resend invitation

Resend Invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resendInvitationRequest** | [**ResendInvitationRequest**](ResendInvitationRequest.md)| ResendInvitationRequest | 

### Return type

 (empty response body)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthControllerRespondToNewPasswordRequiredChallenge

> LoginResult AuthControllerRespondToNewPasswordRequiredChallenge(ctx, respondToNewPasswordRequiredChallengeRequest)

Respond to new password required challenge

Respond to New Password Required Challenge

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**respondToNewPasswordRequiredChallengeRequest** | [**RespondToNewPasswordRequiredChallengeRequest**](RespondToNewPasswordRequiredChallengeRequest.md)| RespondToNewPasswordRequiredChallengeRequest | 

### Return type

[**LoginResult**](LoginResult.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

