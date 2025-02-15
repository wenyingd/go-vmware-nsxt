/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer CertificateAuthority object
type AlbCertificateAuthority struct {
	// It is a reference to an object of type SSLKeyAndCertificate. 
	CaPath string `json:"ca_path,omitempty"`
	// Name of the object.
	Name string `json:"name,omitempty"`
}
