# {{classname}}

All URIs are relative to *https://api.staging.ratio.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1UsersUserIdActivityActivityIdGet**](ActivityApi.md#V1UsersUserIdActivityActivityIdGet) | **Get** /v1/users/{userId}/activity/{activityId} | This returns an Activity item
[**V1UsersUserIdActivityGet**](ActivityApi.md#V1UsersUserIdActivityGet) | **Get** /v1/users/{userId}/activity | This returns a paginated response of Activity items

# **V1UsersUserIdActivityActivityIdGet**
> ActivityItem V1UsersUserIdActivityActivityIdGet(ctx, userId, activityId, ratioClientId, ratioClientSecret)
This returns an Activity item

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User ID | 
  **activityId** | **string**| Activity ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 

### Return type

[**ActivityItem**](ActivityItem.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UsersUserIdActivityGet**
> ActivityItems V1UsersUserIdActivityGet(ctx, userId, ratioClientId, ratioClientSecret, optional)
This returns a paginated response of Activity items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User ID | 
  **ratioClientId** | **string**| Your Ratio Client Identifier | 
  **ratioClientSecret** | **string**| Your Ratio Client Secret | 
 **optional** | ***ActivityApiV1UsersUserIdActivityGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActivityApiV1UsersUserIdActivityGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageToken** | **optional.String**| Query for the next page of activity using the specified token | 
 **pageSize** | **optional.String**| Query for activity with the specified page size | 
 **fromCreateTime** | **optional.String**| Query for activity created on/after the specified time | 
 **toCreateTime** | **optional.String**| Query for activity created on/before the specified time | 
 **walletId** | **optional.String**| Query for activity using the specified wallet ID | 
 **cryptoCurrency** | **optional.String**| Query for activity containing the specified crypto currency | 
 **cryptoStatus** | **optional.String**| Query for activity with crypto in the specified status(es) | 
 **fiatStatus** | **optional.String**| Query for activity with fiat in the specified status(es) | 

### Return type

[**ActivityItems**](ActivityItems.md)

### Authorization

[ClientId](../README.md#ClientId), [ClientSecret](../README.md#ClientSecret)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

