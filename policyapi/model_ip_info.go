/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Only support IP address or subnet. Its type can be of IPv4 or IPv6. It will be converted to subnet when netmask is specified(e.g., 192.168.1.3/24 => 192.168.1.0/24, 2008:12:12:12::2/64 => 2008:12:12:12::/64). 
type IpInfo struct {
	// The destination IP address or subnet
	DstIp string `json:"dst_ip,omitempty"`
	// The source IP address or subnet
	SrcIp string `json:"src_ip,omitempty"`
}