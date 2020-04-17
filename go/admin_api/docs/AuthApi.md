# \AuthApi

All URIs are relative to *https://api.formant.io/v1/admin*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthControllerAdminSignup**](AuthApi.md#AuthControllerAdminSignup) | **Post** /auth/admin-signup | Admin signup
[**AuthControllerChangePassword**](AuthApi.md#AuthControllerChangePassword) | **Post** /auth/change-password | Change password
[**AuthControllerConfirmForgotPassword**](AuthApi.md#AuthControllerConfirmForgotPassword) | **Post** /auth/confirm-forgot-password | Confirm forgot password
[**AuthControllerDeviceCredentials**](AuthApi.md#AuthControllerDeviceCredentials) | **Post** /auth/device-credentials | Device credentials
[**AuthControllerForgotPassword**](AuthApi.md#AuthControllerForgotPassword) | **Post** /auth/forgot-password | Forgot password
[**AuthControllerLogin**](AuthApi.md#AuthControllerLogin) | **Post** /auth/login | Login
[**AuthControllerRefresh**](AuthApi.md#AuthControllerRefresh) | **Post** /auth/refresh | Refresh
[**AuthControllerResendConfirmationCode**](AuthApi.md#AuthControllerResendConfirmationCode) | **Post** /auth/resend-confirmation-code | Resend confirmation code
[**AuthControllerResendInvitation**](AuthApi.md#AuthControllerResendInvitation) | **Post** /auth/resend-invitation | Resend invitation
[**AuthControllerRespondToNewPasswordRequiredChallenge**](AuthApi.md#AuthControllerRespondToNewPasswordRequiredChallenge) | **Post** /auth/respond-to-new-password-required-challenge | Respond to new password required challenge



## AuthControllerAdminSignup

> User AuthControllerAdminSignup(ctx, adminSignupRequest)

Admin signup

Administrator signup

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**adminSignupRequest** | [**AdminSignupRequest**](AdminSignupRequest.md)| AdminSignupRequest | 

### Return type

[**User**](User.md)

### Authorization

Use admin JWT for authorization

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## AuthControllerDeviceCredentials

> DeviceCredentials AuthControllerDeviceCredentials(ctx, )

Device credentials

Device Credentials

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DeviceCredentials**](DeviceCredentials.md)

### Authorization

Use admin JWT for authorization

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

