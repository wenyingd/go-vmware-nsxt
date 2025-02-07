/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Details where NVDS will be migrated to
type ComputeManagerTopology struct {
	// Identifier of vcenter where VDS will be created
	ComputeManagerId string `json:"compute_manager_id"`
	// Datacenter, VDS mapping
	Dvswitch []VdsTopology `json:"dvswitch"`
}
