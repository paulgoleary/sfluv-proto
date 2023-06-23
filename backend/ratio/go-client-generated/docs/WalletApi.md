# {{classname}}

All URIs are relative to *https://api.staging.ratio.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UsersUserIdWalletsGet**](WalletApi.md#V1UsersUserIdWalletsGet) | **Get** /v1/users/{userId}/wallets | This returns the wallets for a user
[**V1UsersUserIdWalletsPost**](WalletApi.md#V1UsersUserIdWalletsPost) | **Post** /v1/users/{userId}/wallets | This connects a wallet to a user
[**V1UsersUserIdWalletsWalletIdGet**](WalletApi.md#V1UsersUserIdWalletsWalletIdGet) | **Get** /v1/users/{userId}/wallets/{walletId} | This returns the specified wallet for a user
[**V1UsersUserIdWalletsWalletIdPatch**](WalletApi.md#V1UsersUserIdWalletsWalletIdPatch) | **Patch** /v1/users/{userId}/wallets/{walletId} | This updates a wallet for a user

# **V1UsersUserIdWalletsGet**
> Wallets V1UsersUserIdWalletsGet(ctx, userId, ratioClientId, ratioClientSecret)
This returns the wallets for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**Wallets**](Wallets.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdWalletsPost**
> Wallet V1UsersUserIdWalletsPost(ctx, body, userId, ratioClientId, ratioClientSecret)
This connects a wallet to a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectWalletRequest**](ConnectWalletRequest.md)| Wallet to be connected | 
  **userId** | **string**| User ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdWalletsWalletIdGet**
> Wallet V1UsersUserIdWalletsWalletIdGet(ctx, userId, walletId, ratioClientId, ratioClientSecret)
This returns the specified wallet for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User ID | 
  **walletId** | **string**| Wallet ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdWalletsWalletIdPatch**
> Wallet V1UsersUserIdWalletsWalletIdPatch(ctx, body, userId, walletId, ratioClientId, ratioClientSecret)
This updates a wallet for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateWalletRequest**](UpdateWalletRequest.md)| Wallet to be updated | 
  **userId** | **string**| User ID | 
  **walletId** | **string**| Wallet ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret), [JWT](../README.md#JWT)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

