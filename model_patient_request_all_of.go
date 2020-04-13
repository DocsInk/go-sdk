/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// PatientRequestAllOf struct for PatientRequestAllOf
type PatientRequestAllOf struct {
	Fname    string `json:"fname,omitempty"`
	Lname    string `json:"lname,omitempty"`
	Address  string `json:"address,omitempty"`
	Address2 string `json:"address2,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Zip      string `json:"zip,omitempty"`
}
