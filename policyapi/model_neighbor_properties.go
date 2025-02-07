/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Neighbor properties
type NeighborProperties struct {
	// Capabilities
	Capabilities string `json:"capabilities,omitempty"`
	// Enabled capabilities
	EnabledCapabilities string `json:"enabled_capabilities,omitempty"`
	// Interface index
	Ifindex int64 `json:"ifindex,omitempty"`
	// Aggregation Capability
	LinkAggregationCapable bool `json:"link_aggregation_capable,omitempty"`
	// Aggregation port id
	LinkAggregationPortId string `json:"link_aggregation_port_id,omitempty"`
	// True if currently in aggregation
	LinkAggregationStatus bool `json:"link_aggregation_status,omitempty"`
	// Interface MAC address
	Mac string `json:"mac,omitempty"`
	// Management address
	MgmtAddr string `json:"mgmt_addr,omitempty"`
	// Interface name
	Name string `json:"name,omitempty"`
	// Object identifier
	Oid string `json:"oid,omitempty"`
	// Port description
	PortDesc string `json:"port_desc,omitempty"`
	// System description
	SystemDesc string `json:"system_desc,omitempty"`
	// System name
	SystemName string `json:"system_name,omitempty"`
	// System port number
	SystemPortNumber int64 `json:"system_port_number,omitempty"`
}
