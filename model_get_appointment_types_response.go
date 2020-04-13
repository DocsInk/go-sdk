/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// GetAppointmentTypesResponse struct for GetAppointmentTypesResponse
type GetAppointmentTypesResponse struct {
	Data  []AppointmentType      `json:"data,omitempty"`
	Links Links                  `json:"links,omitempty"`
	Meta  map[string]interface{} `json:"meta,omitempty"`
}
