# {{classname}}

All URIs are relative to *https://api.staging.ratio.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UsersUserIdBanksBankIdAchPost**](BankApi.md#V1UsersUserIdBanksBankIdAchPost) | **Post** /v1/users/{userId}/banks/{bankId}/ach | This submits a request to initiate an ACH exchange for Crypto
[**V1UsersUserIdBanksBankIdDelete**](BankApi.md#V1UsersUserIdBanksBankIdDelete) | **Delete** /v1/users/{userId}/banks/{bankId} | Deletes a Bank link for a user
[**V1UsersUserIdBanksBankIdGet**](BankApi.md#V1UsersUserIdBanksBankIdGet) | **Get** /v1/users/{userId}/banks/{bankId} | Get a Bank Account by ID for a user
[**V1UsersUserIdBanksBankIdrequestLinkPost**](BankApi.md#V1UsersUserIdBanksBankIdrequestLinkPost) | **Post** /v1/users/{userId}/banks/{bankId}:requestLink | Requests an update link token to re-login to an existing bank account
[**V1UsersUserIdBanksactivateLinkPost**](BankApi.md#V1UsersUserIdBanksactivateLinkPost) | **Post** /v1/users/{userId}/banks:activateLink | Activates a bank account link for a user
[**V1UsersUserIdBanksrequestLinkPost**](BankApi.md#V1UsersUserIdBanksrequestLinkPost) | **Post** /v1/users/{userId}/banks:requestLink | Requests a bank account link token for a user

# **V1UsersUserIdBanksBankIdAchPost**
> InitiateAchResponse V1UsersUserIdBanksBankIdAchPost(ctx, body, userId, bankId, ratioClientId, ratioClientSecret)
This submits a request to initiate an ACH exchange for Crypto

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InitiateAchRequest**](InitiateAchRequest.md)| Request object | 
  **userId** | **string**| User ID | 
  **bankId** | **string**| Bank Account ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**InitiateAchResponse**](InitiateAchResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdBanksBankIdDelete**
> interface{} V1UsersUserIdBanksBankIdDelete(ctx, userId, bankId, ratioClientId, ratioClientSecret)
Deletes a Bank link for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User ID | 
  **bankId** | **string**| Bank Account ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdBanksBankIdGet**
> BankAccount V1UsersUserIdBanksBankIdGet(ctx, userId, bankId, ratioClientId, ratioClientSecret)
Get a Bank Account by ID for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User ID | 
  **bankId** | **string**| Bank Account ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**BankAccount**](BankAccount.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdBanksBankIdrequestLinkPost**
> RequestBankLinkResponse V1UsersUserIdBanksBankIdrequestLinkPost(ctx, body, userId, bankId, ratioClientId, ratioClientSecret)
Requests an update link token to re-login to an existing bank account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestBankUpdateLinkRequest**](RequestBankUpdateLinkRequest.md)| Request object | 
  **userId** | **string**| User ID | 
  **bankId** | **string**| Bank ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**RequestBankLinkResponse**](RequestBankLinkResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdBanksactivateLinkPost**
> ActivateBankLinkResponse V1UsersUserIdBanksactivateLinkPost(ctx, body, userId, ratioClientId, ratioClientSecret)
Activates a bank account link for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ActivateBankLinkRequest**](ActivateBankLinkRequest.md)| Request with public token | 
  **userId** | **string**| User ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**ActivateBankLinkResponse**](ActivateBankLinkResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdBanksrequestLinkPost**
> RequestBankLinkResponse V1UsersUserIdBanksrequestLinkPost(ctx, body, userId, ratioClientId, ratioClientSecret)
Requests a bank account link token for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestBankLinkRequest**](RequestBankLinkRequest.md)| Request object | 
  **userId** | **string**| User ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**RequestBankLinkResponse**](RequestBankLinkResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

