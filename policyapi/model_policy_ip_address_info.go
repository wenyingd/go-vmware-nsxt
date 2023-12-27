/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Used to specify the display name and value of the IPv4Address. 
type PolicyIpAddressInfo struct {
	// Value of the IPv4Address. 
	AddressValue string `json:"address_value"`
	// Display name used to help identify the IPv4Address. 
	DisplayName string `json:"display_name,omitempty"`
	// Next hop used in auto-plumbing of static route. If a value is not provided, static route will not be auto-plumbed. 
	NextHop string `json:"next_hop,omitempty"`
}