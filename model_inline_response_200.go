/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0

 */

package api

// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	Data  []AppointmentType      `json:"data,omitempty"`
	Links Links                  `json:"links,omitempty"`
	Meta  map[string]interface{} `json:"meta,omitempty"`
}
