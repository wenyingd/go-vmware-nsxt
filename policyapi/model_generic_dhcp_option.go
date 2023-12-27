/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Define DHCP options other than option 121.
type GenericDhcpOption struct {
	// Code of the dhcp option.
	Code int64 `json:"code"`
	// Value of the option.
	Values []string `json:"values"`
}