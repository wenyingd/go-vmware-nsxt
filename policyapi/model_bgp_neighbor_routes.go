/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// BGP neighbor learned/advertised route details.
type BgpNeighborRoutes struct {
	// Array of BGP neighbor route details per edge node. 
	EdgeNodeRoutes []RoutesPerTransportNode `json:"edge_node_routes,omitempty"`
	// Array of BGP neighbor route details per edge node. 
	EgdeNodeRoutes []RoutesPerTransportNode `json:"egde_node_routes,omitempty"`
	// Enforcement point policy path
	EnforcementPointPath string `json:"enforcement_point_path,omitempty"`
	// BGP neighbor policy path
	NeighborPath string `json:"neighbor_path,omitempty"`
}
