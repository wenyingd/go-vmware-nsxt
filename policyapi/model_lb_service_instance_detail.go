/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type LbServiceInstanceDetail struct {
	// The display name of the resource which the load balancer instance deploys on. 
	AttachmentDisplayName string `json:"attachment_display_name,omitempty"`
	// The path of the resource which the load balancer instance deploys on. 
	AttachmentPath string `json:"attachment_path,omitempty"`
	// The error message for the load balancer instance. If the instance status is NOT_READY, error message will be attached. 
	ErrorMessage string `json:"error_message,omitempty"`
}
