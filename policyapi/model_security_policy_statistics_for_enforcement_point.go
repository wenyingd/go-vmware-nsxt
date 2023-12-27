/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Aggregate statistics of all the rules in a security policy for a specific enforcement point. 
type SecurityPolicyStatisticsForEnforcementPoint struct {
	// Security Policy statistics for a single container cluster
	ContainerClusterPath string `json:"container_cluster_path,omitempty"`
	// Enforcement point to fetch the statistics from.
	EnforcementPoint string `json:"enforcement_point,omitempty"`
	Statistics *SecurityPolicyStatistics `json:"statistics,omitempty"`
}