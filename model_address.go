/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// Address struct for Address
type Address struct {
	Address  string `json:"address,omitempty"`
	Address2 string `json:"address2,omitempty"`
	City     string `json:"city,omitempty"`
	Dir      string `json:"dir,omitempty"`
	MedRecNo string `json:"med_rec_no,omitempty"`
	State    string `json:"state,omitempty"`
	Zip      string `json:"zip,omitempty"`
}
