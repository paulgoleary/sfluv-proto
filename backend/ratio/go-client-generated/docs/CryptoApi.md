# {{classname}}

All URIs are relative to *https://api.staging.ratio.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CryptoPricesGet**](CryptoApi.md#V1CryptoPricesGet) | **Get** /v1/crypto/prices | Get crypto prices

# **V1CryptoPricesGet**
> GetCryptoPricesResponse V1CryptoPricesGet(ctx, ratioClientId, ratioClientSecret, optional)
Get crypto prices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 
 **optional** | ***CryptoApiV1CryptoPricesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CryptoApiV1CryptoPricesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cryptoCurrencies** | **optional.String**| The crypto currencies to get prices for (comma separated) | 
 **fiatCurrency** | **optional.String**| The fiat currency to quote the prices in | 
 **includeNetworkFees** | **optional.String**| Whether to include network fees in the response | 

### Return type

[**GetCryptoPricesResponse**](GetCryptoPricesResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

