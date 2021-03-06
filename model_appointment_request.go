/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

import (
	"time"
)

// AppointmentRequest struct for AppointmentRequest
type AppointmentRequest struct {
	AppointmentDate   time.Time `json:"appointment_date,omitempty"`
	PatientId         int32     `json:"patient_id,omitempty"`
	AppointmentTypeId int32     `json:"appointment_type_id,omitempty"`
}
