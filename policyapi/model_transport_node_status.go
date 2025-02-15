/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type TransportNodeStatus struct {
	AgentStatus *AgentStatusCount `json:"agent_status,omitempty"`
	ControlConnectionStatus *StatusCount `json:"control_connection_status,omitempty"`
	// Management connection status
	MgmtConnectionStatus string `json:"mgmt_connection_status,omitempty"`
	// Transport node display name
	NodeDisplayName string `json:"node_display_name,omitempty"`
	// Transport node path
	NodePath string `json:"node_path,omitempty"`
	NodeStatus *NodeStatus `json:"node_status,omitempty"`
	// Transport node uuid
	NodeUuid string `json:"node_uuid,omitempty"`
	PnicStatus *StatusCount `json:"pnic_status,omitempty"`
	// Roll-up status of pNIC, management connection, control connection, tunnel status, agent status
	Status string `json:"status,omitempty"`
	ThreatStatus *ThreatStatus `json:"threat_status,omitempty"`
	TunnelStatus *TunnelStatusCount `json:"tunnel_status,omitempty"`
}
