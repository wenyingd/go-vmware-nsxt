/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Describes the capacity and current usage of virtual servers, pools and pool members for the given load balancer service. 
type LbServiceUsage struct {
	Alarm *PolicyRuntimeAlarm `json:"alarm,omitempty"`
	// Policy Path referencing the enforcement point where the info is fetched. 
	EnforcementPointPath string `json:"enforcement_point_path,omitempty"`
	ResourceType string `json:"resource_type"`
	// The current number of pools which has been configured in the given load balancer service. 
	CurrentPoolCount int64 `json:"current_pool_count,omitempty"`
	// The current number of pool members which has been configured in the given load balancer service. 
	CurrentPoolMemberCount int64 `json:"current_pool_member_count,omitempty"`
	// The current number of virtual servers which has been configured in the given load balancer service. 
	CurrentVirtualServerCount int64 `json:"current_virtual_server_count,omitempty"`
	// Timestamp when the data was last updated.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Pool capacity means maximum number of pools which could be configured in the given load balancer service. 
	PoolCapacity int64 `json:"pool_capacity,omitempty"`
	// Pool member capacity means maximum number of pool members which could be configured in the given load balancer service. 
	PoolMemberCapacity int64 `json:"pool_member_capacity,omitempty"`
	// LBService object path.
	ServicePath string `json:"service_path,omitempty"`
	// The size of load balancer service.
	ServiceSize string `json:"service_size,omitempty"`
	// The severity calculation is based on the largest usage percentage from virtual servers, pools and pool members for one load balancer service. 
	Severity string `json:"severity,omitempty"`
	// The usage percentage is the largest usage percentage from virtual servers, pools and pool members for the load balancer service. If the property relax_scale_validation is set as true for LBService, it is possible that the value is larger than 100.0. For example, if SMALL LBS is deployed on MEDIUM edge node and configured with MEDIUM LBS virtual server scale number, LBS usage percentage is shown larger than 100.0. 
	UsagePercentage float32 `json:"usage_percentage,omitempty"`
	// Virtual server capacity means maximum number of virtual servers which could be configured in the given load balancer service. 
	VirtualServerCapacity int64 `json:"virtual_server_capacity,omitempty"`
}