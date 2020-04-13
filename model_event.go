/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// Event struct for Event
type Event struct {
	AppId     int32                  `json:"app_id,omitempty"`
	OrgId     int32                  `json:"org_id,omitempty"`
	EventName string                 `json:"event_name,omitempty"`
	EventTime string                 `json:"event_time,omitempty"`
	Payload   map[string]interface{} `json:"payload,omitempty"`
}
