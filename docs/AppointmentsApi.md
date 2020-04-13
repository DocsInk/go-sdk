# \AppointmentsApi

All URIs are relative to *http://api.docsink.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppointment**](AppointmentsApi.md#CreateAppointment) | **Post** /appointments | Create appointment
[**CreateAppointmentType**](AppointmentsApi.md#CreateAppointmentType) | **Post** /appointments/types | Create appointment types
[**DeleteAppointment**](AppointmentsApi.md#DeleteAppointment) | **Delete** /appointments/{uuid} | Delete appointment
[**GetAppointment**](AppointmentsApi.md#GetAppointment) | **Get** /appointments/{uuid} | Get appointment
[**GetAppointmentTypes**](AppointmentsApi.md#GetAppointmentTypes) | **Get** /appointments/types | Get appointment types
[**UpdateAppointment**](AppointmentsApi.md#UpdateAppointment) | **Put** /appointments/{uuid} | Update Appointment
[**UpdateAppointmentType**](AppointmentsApi.md#UpdateAppointmentType) | **Put** /appointments/types/{uuid} | Update appointment type



## CreateAppointment

> AppointmentResponse CreateAppointment(ctx, optional)

Create appointment

Allows you to create appointments

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateAppointmentOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAppointmentOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appointmentRequest** | [**optional.Interface of AppointmentRequest**](AppointmentRequest.md)|  | 

### Return type

[**AppointmentResponse**](AppointmentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAppointmentType

> AppointmentTypeResponse CreateAppointmentType(ctx, appointmentType)

Create appointment types

create appointment types

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appointmentType** | [**AppointmentType**](AppointmentType.md)|  | 

### Return type

[**AppointmentTypeResponse**](AppointmentTypeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppointment

> CancelResponse DeleteAppointment(ctx, uuid)

Delete appointment

Delete a patient appointment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Appointment UUID | 

### Return type

[**CancelResponse**](CancelResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppointment

> AppointmentResponse GetAppointment(ctx, uuid)

Get appointment

Get appointment

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Appointment id | 

### Return type

[**AppointmentResponse**](AppointmentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppointmentTypes

> GetAppointmentTypesResponse GetAppointmentTypes(ctx, )

Get appointment types

Get all appointment types

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**GetAppointmentTypesResponse**](GetAppointmentTypesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppointment

> AppointmentResponse UpdateAppointment(ctx, uuid, appointmentRequest)

Update Appointment

Update appointment information

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **int32**| Appointment UUID | 
**appointmentRequest** | [**AppointmentRequest**](AppointmentRequest.md)|  | 

### Return type

[**AppointmentResponse**](AppointmentResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppointmentType

> AppointmentTypeResponse UpdateAppointmentType(ctx, uuid, appointmentType)

Update appointment type

update appointment type

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Appointment type UUID | 
**appointmentType** | [**AppointmentType**](AppointmentType.md)|  | 

### Return type

[**AppointmentTypeResponse**](AppointmentTypeResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

