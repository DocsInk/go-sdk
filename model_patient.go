/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// Patient Patient
type Patient struct {
	Uuid      int32    `json:"uuid,omitempty"`
	UserUuid  int32    `json:"user_uuid,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   Address  `json:"address,omitempty"`
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
	DoNotSms  string   `json:"do_not_sms,omitempty"`
}
