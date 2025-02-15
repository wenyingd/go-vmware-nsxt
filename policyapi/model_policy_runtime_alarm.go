/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Alarm associated with the PolicyRuntimeInfoPerEP that exposes potential errors when retrieving runtime information from the enforcement point. 
type PolicyRuntimeAlarm struct {
	ErrorDetails *PolicyApiError `json:"error_details,omitempty"`
	// Alarm error id.
	ErrorId string `json:"error_id,omitempty"`
	// Error message describing the issue.
	Message string `json:"message,omitempty"`
}
