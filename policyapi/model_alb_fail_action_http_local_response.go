/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer FailActionHTTPLocalResponse object
type AlbFailActionHttpLocalResponse struct {
	File *AlbhttpLocalFile `json:"file,omitempty"`
	// Enum options - FAIL_HTTP_STATUS_CODE_200, FAIL_HTTP_STATUS_CODE_503. Default value when not specified in API or module is interpreted by ALB Controller as FAIL_HTTP_STATUS_CODE_503. 
	StatusCode string `json:"status_code,omitempty"`
}
