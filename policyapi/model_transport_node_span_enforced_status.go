/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Detailed Realized Status of an Intent on a span of Transport Nodes. 
type TransportNodeSpanEnforcedStatus struct {
	// Enforced Realized Status Per Scope Resource Type. 
	ResourceType string `json:"resource_type"`
	// List of Detailed Realized Status per Transport Node.
	EnforcedStatusPerTransportNode []EnforcedStatusPerTransportNode `json:"enforced_status_per_transport_node,omitempty"`
}
