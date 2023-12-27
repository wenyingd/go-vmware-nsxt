/*
 * NSX Manager API
 *
 * VMware NSX Manager REST API
 *
 * API version: 4.1.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package containerinventory

import "github.com/vmware/go-vmware-nsxt/common"

// Details of Container Egress.
type ContainerEgress struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []common.ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string                   `json:"_schema,omitempty"`
	Self   *common.SelfResourceLink `json:"_self,omitempty"`
	// Timestamp of last modification
	LastSyncTime int64 `json:"_last_sync_time,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type"`
	// Specifies list of scope of discovered resource. e.g. if VHC path is associated with principal identity, who owns the discovered resource, then scope id will be VHC path and scope type will be VHC.
	Scope []DiscoveredResourceScope `json:"scope,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []common.Tag `json:"tags,omitempty"`
	// Identifier of the container cluster this egress belongs to.
	ContainerClusterId string `json:"container_cluster_id,omitempty"`
	// Container egress IP.
	EgressIp string `json:"egress_ip,omitempty"`
	// Identifier of the container egress.
	ExternalId string `json:"external_id"`
	// Array of additional specific properties of container egress in key-value format.
	OriginProperties []common.KeyValuePair `json:"origin_properties,omitempty"`
	// Container egress specification.
	Spec string `json:"spec,omitempty"`
}