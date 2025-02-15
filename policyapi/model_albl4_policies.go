/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer L4Policies object
type Albl4Policies struct {
	// Index of the virtual service L4 policy set.
	Index int64 `json:"index"`
	// ID of the virtual service L4 policy set. It is a reference to an object of type L4PolicySet. 
	L4PolicySetPath string `json:"l4_policy_set_path"`
}
