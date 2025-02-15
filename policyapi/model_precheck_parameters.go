/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Parameters for nvds upgrade precheck
type PrecheckParameters struct {
	// Cluster ID list for nvds upgrade precheck
	ClusterIds []string `json:"cluster_ids,omitempty"`
}
