/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Used to specify subnets in L3Vpn rule. 
type L3VpnSubnet struct {
	// Subnet used in L3Vpn Rule. 
	Subnet string `json:"subnet"`
}