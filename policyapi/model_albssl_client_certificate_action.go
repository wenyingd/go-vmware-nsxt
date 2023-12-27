/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer SSLClientCertificateAction object
type AlbsslClientCertificateAction struct {
	// Placeholder for description of property close_connection of obj type SSLClientCertificateAction field type str  type boolean. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	CloseConnection bool `json:"close_connection,omitempty"`
	// Placeholder for description of property headers of obj type SSLClientCertificateAction field type str  type array. 
	Headers []AlbsslClientRequestHeader `json:"headers,omitempty"`
}