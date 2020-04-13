# Go API client for api

DocsInk REST API

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./api"
```

## Documentation for API Endpoints

All URIs are relative to *http://api.docsink.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AppointmentsApi* | [**CreateAppointment**](docs/AppointmentsApi.md#createappointment) | **Post** /appointments | Create appointment
*AppointmentsApi* | [**CreateAppointmentType**](docs/AppointmentsApi.md#createappointmenttype) | **Post** /appointments/types | Create appointment types
*AppointmentsApi* | [**DeleteAppointment**](docs/AppointmentsApi.md#deleteappointment) | **Delete** /appointments/{uuid} | Delete appointment
*AppointmentsApi* | [**GetAppointment**](docs/AppointmentsApi.md#getappointment) | **Get** /appointments/{uuid} | Get appointment
*AppointmentsApi* | [**GetAppointmentTypes**](docs/AppointmentsApi.md#getappointmenttypes) | **Get** /appointments/types | Get appointment types
*AppointmentsApi* | [**UpdateAppointment**](docs/AppointmentsApi.md#updateappointment) | **Put** /appointments/{uuid} | Update Appointment
*AppointmentsApi* | [**UpdateAppointmentType**](docs/AppointmentsApi.md#updateappointmenttype) | **Put** /appointments/types/{uuid} | Update appointment type
*PatientsApi* | [**CreatePatient**](docs/PatientsApi.md#createpatient) | **Post** /patients | Create patient
*PatientsApi* | [**GetPatientByMRN**](docs/PatientsApi.md#getpatientbymrn) | **Get** /patients/mrn/{mrn} | Get patient by MRN
*PatientsApi* | [**UpdatePatient**](docs/PatientsApi.md#updatepatient) | **Put** /patients/{uuid} | Update Patient


## Documentation For Models

 - [Address](docs/Address.md)
 - [Appointment](docs/Appointment.md)
 - [AppointmentRequest](docs/AppointmentRequest.md)
 - [AppointmentResponse](docs/AppointmentResponse.md)
 - [AppointmentType](docs/AppointmentType.md)
 - [AppointmentTypeResponse](docs/AppointmentTypeResponse.md)
 - [CancelResponse](docs/CancelResponse.md)
 - [Event](docs/Event.md)
 - [Exception](docs/Exception.md)
 - [GetAppointmentTypesResponse](docs/GetAppointmentTypesResponse.md)
 - [Links](docs/Links.md)
 - [Meta](docs/Meta.md)
 - [Patient](docs/Patient.md)
 - [PatientRequest](docs/PatientRequest.md)
 - [PatientRequestAllOf](docs/PatientRequestAllOf.md)
 - [PatientResponse](docs/PatientResponse.md)


## Documentation For Authorization



## BearerAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```



## Author



