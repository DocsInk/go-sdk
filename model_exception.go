/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// Exception struct for Exception
type Exception struct {
	Exception string `json:"exception,omitempty"`
	File      string `json:"file,omitempty"`
	Line      int32  `json:"line,omitempty"`
	Message   string `json:"message,omitempty"`
}
