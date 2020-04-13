package api

import (
	"context"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
)

var client *APIClient

func TestMain(m *testing.M) {
	var basePath string = "https://dev-api.docsink.com/api"
	var token string = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6ImNlN2Q5YWIzMGMwODk2NjliNmFkYWI2MTY0Y2I4N2U0ZTc4NTc0OTg0ZjFlM2ExYTY1OTYwODBhOWNjYWMxMGExY2NlMTBkNzVhNmUwYTc4In0.eyJhdWQiOiI2OSIsImp0aSI6ImNlN2Q5YWIzMGMwODk2NjliNmFkYWI2MTY0Y2I4N2U0ZTc4NTc0OTg0ZjFlM2ExYTY1OTYwODBhOWNjYWMxMGExY2NlMTBkNzVhNmUwYTc4IiwiaWF0IjoxNTgxODAxNDU3LCJuYmYiOjE1ODE4MDE0NTcsImV4cCI6MjUyODU3MjY1Nywic3ViIjoiIiwic2NvcGVzIjpbXX0.xJYkdQr5fnMFih6D1JdaVRtcnddtGbDYOBBEhZjxFcCVo8AE4YrClmQZDYDz-RFwSeWA4vJj_221THlG5qOO8nWE_831GhhlaYi-lNQmBUa-Tewi4CoOtfzy4ku3TZY6mljd46tswjIzmpmkPcuNmTFgqsOazzGIeGtMsoD-RihmUWrUpC_buBRpxHMNDy7rz7rxyozeWMJ61GyU14UtxIFa71-4KifW3yeOcnWYS2kS1KvHvoS1yejA3Zsz9cN7qLghfLJaHXEIw8mfl0_H0Yo6uNH74equtY_Q94Jm8Po-2-ykTQiGXtxSV78L25JxMX-FZzk82shlR0Rhkp2JUMvOSvNRSuAtL2op-Gm9CbzS_U1V17FgMgjs_qjo8l3JOl5zZOGAjZcdqbX9VBG4anefKiv0BXJMUT1n9kjQ29b9bqiAa0FRp7Wap9u8D7-S6Q1f7HUK8P41BVP0Fr1AYEzstfbHKT7dEjhP5oNov_tC1Xip-YzIO6uow0XgVd1V2sSHVTDdgD_G0A9We0INSCq7Y3D0xDBr9DT16x4e3BDDwc0RywIHFw1dqsGC5ytS9TG7oJ-OpkHn2j_BRfKYX8BOFV95Qu8nlPpI2xWizOjzDcCdpu0vUh6I7P3TD_wT0D1pRkjARDVHyswPuLgHGHATP2n2_XqDLOcjvteRymU"

	config := NewConfiguration()
	config.AddDefaultHeader("Authorization", token)

	client = NewAPIClient(config)
	client.ChangeBasePath(basePath)

	m.Run()
}

func TestCreatePatient(t *testing.T) {
	gender := randomdata.RandomGender
	dob, _ := time.Parse("Monday 2 Jan 2006", randomdata.FullDate())
	admitDate, _ := time.Parse("Monday 2 Jan 2006", randomdata.FullDate())

	patient := PatientRequest{}
	patient.Fname = randomdata.FirstName(gender)
	patient.Lname = randomdata.LastName()
	patient.Email = randomdata.Email()
	patient.PhoneNo = randomdata.Digits(10)
	patient.Dob = dob.Format("2006-01-02")
	patient.AdmitDate = admitDate.Format("2006-01-02")

	if gender == 0 {
		patient.Gender = "M"
	} else {
		patient.Gender = "F"
	}

	patient.MedRecNo = randomdata.Digits(8)

	resp, _, err := client.PatientsApi.CreatePatient(context.Background(), patient)
	if err != nil {
		t.Error(err)
	}

	if resp.Data.Uuid == 0 {
		t.Errorf("Patient UUID not set")
	}
}

func TestCreateAppointment(t *testing.T) {
	// apptTypes, _, err := client.AppointmentsApi.GetAppointmentTypes(context.Background())
	// if err != nil {
	// 	t.Error(err)
	// }
}

func TestUpdateAppointment(t *testing.T) {
	// appt := Appointment{}
	// patient := Patient{}

	// patientRequest = Patient{
	// 	Fname: randomdata.FirstName(gender),
	// 	Lname: randomdata.LastName(),
	// 	// Address: appointment.Patient.Address1,
	// 	Email:     randomdata.Email(),
	// 	PhoneNo:   randomdata.Digits(10),
	// 	Dob:       dob.Format("2006-01-02"),
	// 	AdmitDate: admitDate.Format("2006-01-02"),
	// 	MedRecNo:  athenaPatient.Patientid,
	// 	Census:    false,
	// }

	// appointmentDateTime, err := time.Parse(time.RFC3339, "2020-02-03T01:00:00-05:00")
	// if err != nil {
	// 	t.Error(err)
	// }

}

func TestGetAppointmentTypes(t *testing.T) {
	_, _, err := client.AppointmentsApi.GetAppointmentTypes(context.Background())
	if err != nil {
		t.Error(err)
	}
}
