/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HSMThalesRFS object
type AlbhsmThalesRfs struct {
	Ip *AlbIpAddr `json:"ip"`
	// Port at which the RFS server accepts the sync request from clients for Thales encrypted private key. Allowed values are 1-65535. Default value when not specified in API or module is interpreted by ALB Controller as 9004. 
	Port int64 `json:"port,omitempty"`
}
