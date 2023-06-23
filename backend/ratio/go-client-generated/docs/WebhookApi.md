# {{classname}}

All URIs are relative to *https://api.staging.ratio.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1WebhooksEventsGet**](WebhookApi.md#V1WebhooksEventsGet) | **Get** /v1/webhooks/events | Get all webhook events
[**V1WebhooksGet**](WebhookApi.md#V1WebhooksGet) | **Get** /v1/webhooks | This returns the client&#x27;s webhooks
[**V1WebhooksPost**](WebhookApi.md#V1WebhooksPost) | **Post** /v1/webhooks | This creates a new webhook
[**V1WebhooksWebhookIdDelete**](WebhookApi.md#V1WebhooksWebhookIdDelete) | **Delete** /v1/webhooks/{webhookId} | This deletes the requested webhook
[**V1WebhooksWebhookIdGet**](WebhookApi.md#V1WebhooksWebhookIdGet) | **Get** /v1/webhooks/{webhookId} | This returns the requested webhook
[**V1WebhooksWebhookIdPatch**](WebhookApi.md#V1WebhooksWebhookIdPatch) | **Patch** /v1/webhooks/{webhookId} | This updates the requested webhook

# **V1WebhooksEventsGet**
> WebhookEvents V1WebhooksEventsGet(ctx, ratioClientId, ratioClientSecret)
Get all webhook events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**WebhookEvents**](WebhookEvents.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WebhooksGet**
> Webhooks V1WebhooksGet(ctx, ratioClientId, ratioClientSecret)
This returns the client's webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**Webhooks**](Webhooks.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WebhooksPost**
> Webhook V1WebhooksPost(ctx, body, ratioClientId, ratioClientSecret)
This creates a new webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateWebhookRequest**](CreateWebhookRequest.md)| Request object | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WebhooksWebhookIdDelete**
> interface{} V1WebhooksWebhookIdDelete(ctx, webhookId, ratioClientId, ratioClientSecret)
This deletes the requested webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **string**| The id of the webhook | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WebhooksWebhookIdGet**
> Webhook V1WebhooksWebhookIdGet(ctx, webhookId, ratioClientId, ratioClientSecret)
This returns the requested webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **string**| The id of the webhook | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1WebhooksWebhookIdPatch**
> Webhook V1WebhooksWebhookIdPatch(ctx, body, webhookId, ratioClientId, ratioClientSecret)
This updates the requested webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md)| Request object | 
  **webhookId** | **string**| The id of the webhook | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

