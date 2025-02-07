/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// The current runtime status of the DNS forwarder. 
type NsxTdnsForwarderStatus struct {
	// Policy path referencing the enforcement point from where the status is fetched. 
	EnforcementPointPath string `json:"enforcement_point_path,omitempty"`
	ResourceType string `json:"resource_type"`
	// Extra message, if available
	ExtraMessage string `json:"extra_message,omitempty"`
	// UP means the DNS forwarder is working correctly on the active transport node and the stand-by transport node (if present). Failover will occur if either node goes down. DOWN means the DNS forwarder is down on both active transport node and standby node (if present). The DNS forwarder does not function in this situation. Error means there is some error on one or both transport node, or no status was reported from one or both transport nodes. The DNS forwarder may be working (or not working). NO_BACKUP means DNS forwarder is working in only one transport node, either because it is down on the standby node, or no standby is configured. An forwarder outage will occur if the active node goes down. 
	Status string `json:"status,omitempty"`
	// Time stamp of the current status, in ms
	Timestamp int64 `json:"timestamp,omitempty"`
}
