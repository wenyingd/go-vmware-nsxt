/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type LbServiceStatistics struct {
	Alarm *PolicyRuntimeAlarm `json:"alarm,omitempty"`
	// Policy Path referencing the enforcement point where the info is fetched. 
	EnforcementPointPath string `json:"enforcement_point_path,omitempty"`
	ResourceType string `json:"resource_type"`
	// Timestamp when the data was last updated.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Statistics of load balancer pools
	Pools []LbPoolStatistics `json:"pools,omitempty"`
	// load balancer service identifier.
	ServicePath string `json:"service_path,omitempty"`
	Statistics *LbServiceStatisticsCounter `json:"statistics,omitempty"`
	// Statistics of load balancer virtual servers.
	VirtualServers []LbVirtualServerStatistics `json:"virtual_servers,omitempty"`
}
