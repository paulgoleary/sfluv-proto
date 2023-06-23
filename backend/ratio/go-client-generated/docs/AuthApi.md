# {{classname}}

All URIs are relative to *https://api.staging.ratio.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AuthCryptoWalletaddToUserPost**](AuthApi.md#V1AuthCryptoWalletaddToUserPost) | **Post** /v1/auth/cryptoWallet:addToUser | Adds an unverified crypto wallet to a user. A crypto wallet remains unverified until the user authenticates it.
[**V1AuthCryptoWalletauthenticatePost**](AuthApi.md#V1AuthCryptoWalletauthenticatePost) | **Post** /v1/auth/cryptoWallet:authenticate | Authenticate a user&#x27;s crypto wallet
[**V1AuthCryptoWalletstartPost**](AuthApi.md#V1AuthCryptoWalletstartPost) | **Post** /v1/auth/cryptoWallet:start | Start a crypto wallet authentication flow
[**V1AuthMagicLinkverifyPost**](AuthApi.md#V1AuthMagicLinkverifyPost) | **Post** /v1/auth/magicLink:verify | Verify a user&#x27;s email with a magic link
[**V1AuthOtpEmailauthenticatePost**](AuthApi.md#V1AuthOtpEmailauthenticatePost) | **Post** /v1/auth/otp/email:authenticate | Authenticate a user with an Email OTP
[**V1AuthOtpEmailsendPost**](AuthApi.md#V1AuthOtpEmailsendPost) | **Post** /v1/auth/otp/email:send | Send an Email OTP to the user
[**V1AuthOtpSmsauthenticatePost**](AuthApi.md#V1AuthOtpSmsauthenticatePost) | **Post** /v1/auth/otp/sms:authenticate | Authenticate a user with an SMS OTP
[**V1AuthOtpSmssendPost**](AuthApi.md#V1AuthOtpSmssendPost) | **Post** /v1/auth/otp/sms:send | Send an SMS OTP to the user

# **V1AuthCryptoWalletaddToUserPost**
> V1AuthCryptoWalletaddToUserPost(ctx, body, ratioClientId, ratioClientSecret)
Adds an unverified crypto wallet to a user. A crypto wallet remains unverified until the user authenticates it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddUnverifiedCryptoWalletRequest**](AddUnverifiedCryptoWalletRequest.md)| Add Unverified Crypto wallet auth request object | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

 (empty response body)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthCryptoWalletauthenticatePost**
> AuthResponse V1AuthCryptoWalletauthenticatePost(ctx, body, ratioClientId, ratioClientSecret)
Authenticate a user's crypto wallet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthenticateCryptoWalletRequest**](AuthenticateCryptoWalletRequest.md)| Crypto wallet auth request object | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthCryptoWalletstartPost**
> AuthenticateCryptoWalletStartResponse V1AuthCryptoWalletstartPost(ctx, body, ratioClientId, ratioClientSecret)
Start a crypto wallet authentication flow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthenticateCryptoWalletStartRequest**](AuthenticateCryptoWalletStartRequest.md)| Start Crypto wallet auth request object | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**AuthenticateCryptoWalletStartResponse**](AuthenticateCryptoWalletStartResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthMagicLinkverifyPost**
> V1AuthMagicLinkverifyPost(ctx, body)
Verify a user's email with a magic link

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VerifyMagicLinkRequest**](VerifyMagicLinkRequest.md)| Authenticate Email OTP request object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthOtpEmailauthenticatePost**
> AuthResponse V1AuthOtpEmailauthenticatePost(ctx, body, ratioClientId, ratioClientSecret)
Authenticate a user with an Email OTP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthenticateEmailOtpRequest**](AuthenticateEmailOtpRequest.md)| Authenticate Email OTP request object | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthOtpEmailsendPost**
> SendEmailOtpResponse V1AuthOtpEmailsendPost(ctx, body, ratioClientId, ratioClientSecret)
Send an Email OTP to the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendEmailOtpRequest**](SendEmailOtpRequest.md)| Send Email OTP request object | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**SendEmailOtpResponse**](SendEmailOtpResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthOtpSmsauthenticatePost**
> AuthResponse V1AuthOtpSmsauthenticatePost(ctx, body, ratioClientId, ratioClientSecret)
Authenticate a user with an SMS OTP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthenticateSmsOtpRequest**](AuthenticateSmsOtpRequest.md)| Authenticate SMS OTP request object | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1AuthOtpSmssendPost**
> SendSmsOtpResponse V1AuthOtpSmssendPost(ctx, body, ratioClientId, ratioClientSecret)
Send an SMS OTP to the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendSmsOtpRequest**](SendSmsOtpRequest.md)| Send SMS OTP request object | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**SendSmsOtpResponse**](SendSmsOtpResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

