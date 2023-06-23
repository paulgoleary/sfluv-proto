# {{classname}}

All URIs are relative to *https://api.staging.ratio.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ClientSessionsPost**](ClientApi.md#V1ClientSessionsPost) | **Post** /v1/client/sessions | Creates a client session
[**V1ClientSessionsSessionIdGet**](ClientApi.md#V1ClientSessionsSessionIdGet) | **Get** /v1/client/sessions/{sessionId} | Returns the client session

# **V1ClientSessionsPost**
> ClientSession V1ClientSessionsPost(ctx, body, ratioClientId, ratioClientSecret)
Creates a client session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateClientSessionRequest**](CreateClientSessionRequest.md)| Create client session request object | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**ClientSession**](ClientSession.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ClientSessionsSessionIdGet**
> ClientSession V1ClientSessionsSessionIdGet(ctx, sessionId)
Returns the client session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**| Session ID | 

### Return type

[**ClientSession**](ClientSession.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

