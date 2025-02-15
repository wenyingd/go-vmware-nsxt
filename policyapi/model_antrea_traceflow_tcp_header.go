/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Tcp header stuffs for Antrea traceflow. 
type AntreaTraceflowTcpHeader struct {
	// Destination port number in TcpHeader. 
	DstPort int64 `json:"dstPort,omitempty"`
	// Source port number in TcpHeader. 
	SrcPort int64 `json:"srcPort,omitempty"`
	// Tcp flags in TcpHeader. SYN flag must be set for traceflow. 
	TcpFlags int64 `json:"tcpFlags,omitempty"`
}
