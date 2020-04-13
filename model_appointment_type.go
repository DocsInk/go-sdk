/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// AppointmentType struct for AppointmentType
type AppointmentType struct {
	Active      int32  `json:"active,omitempty"`
	Description string `json:"description,omitempty"`
	IsGlobal    int32  `json:"is_global,omitempty"`
	Name        string `json:"name,omitempty"`
	Uuid        int32  `json:"uuid,omitempty"`
}
