/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0

 */

package api

// InlineResponse500 apiResponse
type InlineResponse500 struct {
	Message   int32  `json:"message,omitempty"`
	Exception string `json:"exception,omitempty"`
}
