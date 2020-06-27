# \PatientsApi

All URIs are relative to *http://dev-api.docsink.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePatient**](PatientsApi.md#CreatePatient) | **Post** /patients | Create patient
[**GetPatientByMRN**](PatientsApi.md#GetPatientByMRN) | **Get** /patients/mrn/{mrn} | Get patient by MRN
[**UpdatePatient**](PatientsApi.md#UpdatePatient) | **Put** /patients/{uuid} | Update Patient



## CreatePatient

> CreatePatientResponse CreatePatient(ctx, patientRequest)

Create patient

Create a new DocsInk patient

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patientRequest** | [**PatientRequest**](PatientRequest.md)|  | 

### Return type

[**CreatePatientResponse**](CreatePatientResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPatientByMRN

> PatientResponse GetPatientByMRN(ctx, mrn)

Get patient by MRN

Get patient using external system MRN

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mrn** | **string**| Patient MRN | 

### Return type

[**PatientResponse**](PatientResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePatient

> PatientResponse UpdatePatient(ctx, uuid, patientRequest)

Update Patient

Update patient information

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string**| Patient UUID | 
**patientRequest** | [**PatientRequest**](PatientRequest.md)|  | 

### Return type

[**PatientResponse**](PatientResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

