/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HealthMonitorImap object
type AlbHealthMonitorImap struct {
	// Folder to access.
	Folder string `json:"folder,omitempty"`
	SslAttributes *AlbHealthMonitorSslAttributes `json:"ssl_attributes,omitempty"`
}
