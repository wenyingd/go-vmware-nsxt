/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer NetworkSecurityPolicyActionRLParam object
type AlbNetworkSecurityPolicyActionRlParam struct {
	// Maximum number of connections or requests or packets to be rate limited instantaneously. Default value when not specified in API or module is interpreted by ALB Controller as 0. 
	BurstSize int64 `json:"burst_size"`
	// Maximum number of connections or requests or packets per second. Allowed values are 1-4294967295. 
	MaxRate int64 `json:"max_rate"`
}
