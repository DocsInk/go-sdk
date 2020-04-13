/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// PatientRequest struct for PatientRequest
type PatientRequest struct {
	Uuid      int32    `json:"uuid,omitempty"`
	UserUuid  int32    `json:"user_uuid,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   string   `json:"address,omitempty"`
	Email     string   `json:"email,omitempty"`
	PhoneNo   string   `json:"phone_no,omitempty"`
	Dob       string   `json:"dob,omitempty"`
	AdmitDate string   `json:"admit_date,omitempty"`
	Gender    string   `json:"gender,omitempty"`
	MedRecNo  string   `json:"med_rec_no,omitempty"`
	Census    bool     `json:"census,omitempty"`
	Provider  int32    `json:"provider,omitempty"`
	Location  int32    `json:"location,omitempty"`
	Flags     []string `json:"flags,omitempty"`
	Fname     string   `json:"fname,omitempty"`
	Lname     string   `json:"lname,omitempty"`
	Address2  string   `json:"address2,omitempty"`
	City      string   `json:"city,omitempty"`
	State     string   `json:"state,omitempty"`
	Zip       string   `json:"zip,omitempty"`
}
