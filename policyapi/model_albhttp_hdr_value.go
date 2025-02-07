/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HTTPHdrValue object
type AlbhttpHdrValue struct {
	// HTTP header value or variable representing an HTTP header.
	Val string `json:"val,omitempty"`
	// Variable. Enum options - HTTP_POLICY_VAR_CLIENT_IP, HTTP_POLICY_VAR_VS_PORT, HTTP_POLICY_VAR_VS_IP, HTTP_POLICY_VAR_HTTP_HDR, HTTP_POLICY_VAR_SSL_CLIENT_FINGERPRINT, HTTP_POLICY_VAR_SSL_CLIENT_SERIAL, HTTP_POLICY_VAR_SSL_CLIENT_ISSUER, HTTP_POLICY_VAR_SSL_CLIENT_SUBJECT, HTTP_POLICY_VAR_SSL_CLIENT_RAW, HTTP_POLICY_VAR_SSL_PROTOCOL, HTTP_POLICY_VAR_SSL_SERVER_NAME, HTTP_POLICY_VAR_USER_NAME, HTTP_POLICY_VAR_SSL_CIPHER, HTTP_POLICY_VAR_REQUEST_ID, HTTP_POLICY_VAR_SSL_CLIENT_VERSION, HTTP_POLICY_VAR_SSL_CLIENT_SIGALG, HTTP_POLICY_VAR_SSL_CLIENT_NOTVALIDBEFORE, HTTP_POLICY_VAR_SSL_CLIENT_NOTVALIDAFTER. 
	Var_ string `json:"var,omitempty"`
}
