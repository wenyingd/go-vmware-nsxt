/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// This action is used to reject HTTP request messages. The specified reply_status value is used as the status code for the corresponding HTTP response message which is sent back to client (Normally a browser) indicating the reason it was rejected. Reference official HTTP status code list for your specific HTTP version to set the reply_status properly. LBHttpRejectAction does not support variables. 
type LbHttpRejectAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Response message.
	ReplyMessage string `json:"reply_message,omitempty"`
	// HTTP response status code.
	ReplyStatus string `json:"reply_status"`
}
