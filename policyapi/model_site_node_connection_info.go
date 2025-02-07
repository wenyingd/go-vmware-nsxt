/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Credential info to connect to a node in the federated remote site.
type SiteNodeConnectionInfo struct {
	// Please specify the fqdn of the Management Node of your site. 
	Fqdn string `json:"fqdn"`
	// Password to connect to Site's Local Manager.
	Password string `json:"password,omitempty"`
	// Thumbprint of Site's Local Manager in the form of a SHA-256 hash represented in lower case HEX. 
	Thumbprint string `json:"thumbprint,omitempty"`
	// Username to connect to Site's Local Manager.
	Username string `json:"username,omitempty"`
}
