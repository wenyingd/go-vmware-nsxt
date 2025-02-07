/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Traffic statistics for IPSec VPN Ike session. Note - Not supported in this release. 
type IpSecVpnIkeTrafficStatistics struct {
	// Number of bytes in.
	BytesIn int64 `json:"bytes_in,omitempty"`
	// Number of bytes out.
	BytesOut int64 `json:"bytes_out,omitempty"`
	// Fail count.
	FailCount int64 `json:"fail_count,omitempty"`
	// Number of packets in.
	PacketsIn int64 `json:"packets_in,omitempty"`
	// Number of packets out.
	PacketsOut int64 `json:"packets_out,omitempty"`
}
