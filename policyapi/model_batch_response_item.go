/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// A single respose in a list of batched responses
type BatchResponseItem struct {
	// object returned by api
	Body interface{} `json:"body,omitempty"`
	// http status code
	Code int64 `json:"code"`
	// The headers returned by the API call
	Headers interface{} `json:"headers,omitempty"`
}