/*
 * NSX Manager API
 *
 * VMware NSX Manager REST API
 *
 * API version: 4.1.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package containerinventory

// Specifies an address that can be bound to a container gateway.
type ContainerGatewayAddress struct {
	// Specifies address of container gateway
	Address string `json:"address,omitempty"`
	// Specifies type of gateway address. e.g. Hostname, IPAddress, NamedAddress.
	Type string `json:"type,omitempty"`
}