/*
 * NSX Manager API
 *
 * VMware NSX Manager REST API
 *
 * API version: 4.1.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package containerinventory

// It represents the status of a load-balancer ingress point.
type ContainerLoadBalancerIngress struct {
	// Hostname is set for load-balancer ingress points that are DNS based.
	Hostname string `json:"hostname,omitempty"`
	// IP is set for load-balancer ingress points that are IP based.
	Ip string `json:"ip,omitempty"`
	// Ports is a list of records of service points.
	Ports []PortStatus `json:"ports,omitempty"`
}