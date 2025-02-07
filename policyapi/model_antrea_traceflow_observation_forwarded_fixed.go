/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type AntreaTraceflowObservationForwarded struct {
	// The type of component. 
	ComponentType string `json:"component_type,omitempty"`
	// UID of the container node that observed a traceflow packet. 
	ContainerNodeId string `json:"container_node_id,omitempty"`
	// The type of observation. AntreaTraceflowObservationDelivered: The packet was delivered to destination Pod properly AntreaTraceflowObservationReceived: The packet was received from another ContainerNode AntreaTraceflowObservationForwarded: The packet was forwarded to next logical node or ContainerNode AntreaTraceflowObservationDropped: The packet was dropped 
	ObservationType string `json:"observation_type"`
	// Timestamp when the observation was collect by Antrea controller. 
	Timestamp int64 `json:"timestamp,omitempty"`
	// ID of the Container application instance from which the traceflow packet was forwarded. 
	ContainerApplicationInstanceId string `json:"container_application_instance_id,omitempty"`
	// ID of the hit network policy that makes the traceflow packet forwarded. 
	ContainerNetworkPolicyId string `json:"container_network_policy_id,omitempty"`
	// ID of the hit DFW rule that traceflow packet was forwarded. 
	RuleId string `json:"rule_id,omitempty"`
	// Time to live (TTL) or hop limit of a traceflow packet. 
	Ttl int32 `json:"ttl,omitempty"`
	// The translated source IP/IPv6 address of VPN/NAT. 
	TranslatedSrcIp string `json:"translated_src_ip,omitempty"`
	// The translated destination IP/IPv6 address of VPN/NAT. 
	TranslatedDstIp string `json:"translated_dst_ip,omitempty"`
	// Destination mac address of traceflow packet. 
	DestinationMac string `json:"destination_mac,omitempty"`
	// The tunnel destination IP/IPv6 address. 
	TunnelDstIp string `json:"tunnel_dst_ip,omitempty"`
	// The name of component that packets passed through. 
	ComponentName string `json:"component_name,omitempty"`
}
