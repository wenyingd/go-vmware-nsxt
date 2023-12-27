/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// IDS event flow data specific to each IDS event. The data includes source ip, source port, destination ip, destination port, protocol, rule id, profile id, and the action. 
type PolicyIdsEventFlowData struct {
	// The action pertaining to the detected intrusion. Possible values are ALERT, DROP, REJECT, and INVALID. ALERT - If there is a signature match on the packet, it is allowed to pass but a notification is sent to the user notifying an intrusion was detected. DROP - On a signature match, the packet is silently dropped. An alert is sent to the user that an intrusion was detected. REJECT - On a signature match, the packet is dropped and TCP RST or ICMP error messages (for non-TCP pkts) are sent to the endpoints. An alert is sent to the user that an intrusion was detected. INVALID - If the action doesn't belong to any of the above mentioned categories, it is marked as INVALID.
	ActionType string `json:"action_type,omitempty"`
	// Bytes sent to client.
	BytesToclient int64 `json:"bytes_toclient,omitempty"`
	// Bytes sent to server.
	BytesToserver int64 `json:"bytes_toserver,omitempty"`
	// IP address of the VM that initiated the communication.
	ClientIp string `json:"client_ip,omitempty"`
	// IP address of the destination VM on the intrusion flow.
	DestinationIp string `json:"destination_ip,omitempty"`
	// Port on the destination VM where the traffic was sent to.
	DestinationPort int64 `json:"destination_port,omitempty"`
	// Name of the gateway on which this intrusion was detected.
	Gateway string `json:"gateway,omitempty"`
	// Tags associated with the gateway on which this intrusion was detected.
	GatewayTags []Tag `json:"gateway_tags,omitempty"`
	// Name of the host on which this intrusion was detected.
	Host string `json:"host,omitempty"`
	// IP address of VM on the host where IDS engine is running.
	LocalVmIp string `json:"local_vm_ip,omitempty"`
	// The IDS profile id that is associated with the IDS rule pertaining to the intrusion event detected.
	ProfileId string `json:"profile_id,omitempty"`
	// Traffic protocol pertaining to the detected intrusion, could be TCP/UDP etc.
	Protocol string `json:"protocol,omitempty"`
	// The IDS Rule id pertaining to the detected intrusion.
	RuleId int64 `json:"rule_id,omitempty"`
	// IP address of the source VM on the intrusion flow.
	SourceIp string `json:"source_ip,omitempty"`
	// Source port through which traffic was initiated that caused the intrusion to be detected.
	SourcePort int64 `json:"source_port,omitempty"`
	// The source where the intrusion was detected. Possible values are GATEWAY and HOST.
	TrafficType string `json:"traffic_type,omitempty"`
}