/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Traffic counters for IPSec VPN session.
type IpSecVpnTrafficCounters struct {
	// Total number of bytes recevied.
	BytesIn int64 `json:"bytes_in,omitempty"`
	// Total number of bytes sent.
	BytesOut int64 `json:"bytes_out,omitempty"`
	// Total number of incoming packets dropped on inbound security association. 
	DroppedPacketsIn int64 `json:"dropped_packets_in,omitempty"`
	// Total number of outgoing packets dropped on outbound security association. 
	DroppedPacketsOut int64 `json:"dropped_packets_out,omitempty"`
	// Total number of packets received.
	PacketsIn int64 `json:"packets_in,omitempty"`
	// Total number of packets sent.
	PacketsOut int64 `json:"packets_out,omitempty"`
}