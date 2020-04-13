/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// Appointment Patient appointment
type Appointment struct {
	AppointmentDate     string                 `json:"appointment_date,omitempty"`
	AppointmentDateUtc  string                 `json:"appointment_date_utc,omitempty"`
	AppointmentTimezone string                 `json:"appointment_timezone,omitempty"`
	CreatedAt           string                 `json:"created_at,omitempty"`
	CreatedBy           int32                  `json:"created_by,omitempty"`
	CreatedByFirstName  string                 `json:"created_by_first_name,omitempty"`
	CreatedByLastName   string                 `json:"created_by_last_name,omitempty"`
	InitialSmsSent      bool                   `json:"initial_sms_sent,omitempty"`
	Location            map[string]interface{} `json:"location,omitempty"`
	Patient             Patient                `json:"patient,omitempty"`
	Type                AppointmentType        `json:"type,omitempty"`
	UpdatedAt           string                 `json:"updated_at,omitempty"`
	User                []string               `json:"user,omitempty"`
	Uuid                int32                  `json:"uuid,omitempty"`
}
