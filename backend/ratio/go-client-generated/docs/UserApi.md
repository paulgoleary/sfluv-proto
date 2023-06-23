# {{classname}}

All URIs are relative to *https://api.staging.ratio.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UsersPost**](UserApi.md#V1UsersPost) | **Post** /v1/users | This creates a new user
[**V1UsersUserIdGet**](UserApi.md#V1UsersUserIdGet) | **Get** /v1/users/{userId} | This returns the current authenticated user
[**V1UsersUserIdIdvPost**](UserApi.md#V1UsersUserIdIdvPost) | **Post** /v1/users/{userId}/idv | This initiates the IDV process for user
[**V1UsersUserIdKycPost**](UserApi.md#V1UsersUserIdKycPost) | **Post** /v1/users/{userId}/kyc | This initiates the KYC process for a user
[**V1UsersUserIdcalculateAchLimitsPost**](UserApi.md#V1UsersUserIdcalculateAchLimitsPost) | **Post** /v1/users/{userId}:calculateAchLimits | Calculate the ACH limits for the user

# **V1UsersPost**
> User V1UsersPost(ctx, body, ratioClientId, ratioClientSecret)
This creates a new user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserRequest**](CreateUserRequest.md)| User to be created | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**User**](User.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdGet**
> User V1UsersUserIdGet(ctx, userId, ratioClientId, ratioClientSecret)
This returns the current authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| The id of the user | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**User**](User.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdIdvPost**
> IdvResponse V1UsersUserIdIdvPost(ctx, userId, ratioClientId, ratioClientSecret)
This initiates the IDV process for user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**IdvResponse**](IdvResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdKycPost**
> User V1UsersUserIdKycPost(ctx, body, userId, ratioClientId, ratioClientSecret)
This initiates the KYC process for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitKycRequest**](SubmitKycRequest.md)| Submit kyc request body | 
  **userId** | **string**| User ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**User**](User.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdcalculateAchLimitsPost**
> CalculateAchLimitsResponse V1UsersUserIdcalculateAchLimitsPost(ctx, userId, ratioClientId)
Calculate the ACH limits for the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 

### Return type

[**CalculateAchLimitsResponse**](CalculateAchLimitsResponse.md)

### Authorization

[JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

