/*
 * NSX Manager API
 *
 * VMware NSX Manager REST API
 *
 * API version: 4.1.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package containerinventory

// Specifies details of endpoints, when container application is NodePortLocal
type NodePortLocalEndpoint struct {
	// Specifies IP address of node.
	NodeIp string `json:"node_ip,omitempty"`
	// Specifies the port on the node that can be used to access the service.
	NodePort int64 `json:"node_port,omitempty"`
	// Specifies protocol of endpoint. e.g. TCP, UDP.
	Protocol string `json:"protocol,omitempty"`
}
