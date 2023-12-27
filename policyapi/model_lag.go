/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// LACP group
type Lag struct {
	// unique id
	Id string `json:"id,omitempty"`
	// LACP load balance Algorithm
	LoadBalanceAlgorithm string `json:"load_balance_algorithm"`
	// LACP group mode
	Mode string `json:"mode"`
	// Lag name
	Name string `json:"name"`
	// number of uplinks
	NumberOfUplinks int32 `json:"number_of_uplinks"`
	// LACP timeout type
	TimeoutType string `json:"timeout_type,omitempty"`
	// uplink names
	Uplinks []Uplink `json:"uplinks,omitempty"`
}