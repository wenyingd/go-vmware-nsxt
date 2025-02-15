/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Index for cross site allocation for edge cluster and its members referred by gateway. 
type SiteAllocationIndexForEdge struct {
	// Unqiue edge cluster node index across sites based on stretch of the Gateway. For example, if a Gateway is streched to sites S1 with one edge cluster of 3 nodes and site S2 with one edge cluster of 2 nodes, the in the Global Manager will allocate the index for 5 edge nodes and 2 cluster in the rage 0 to 7. 
	Index int64 `json:"index,omitempty"`
	// Edge cluster or edge node path
	TargetResourcePath string `json:"target_resource_path,omitempty"`
}
