# {{classname}}

All URIs are relative to *https://api.staging.ratio.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ConfigGet**](ConfigApi.md#V1ConfigGet) | **Get** /v1/config | Gets the client configuration

# **V1ConfigGet**
> GetConfigResponse V1ConfigGet(ctx, ratioClientId, ratioClientSecret)
Gets the client configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**GetConfigResponse**](GetConfigResponse.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

