/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Contains BGP routing configuration. 
type BgpRoutingConfig struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// Path of its parent
	ParentPath string `json:"parent_path,omitempty"`
	// Absolute path of this object
	Path string `json:"path,omitempty"`
	// This is a UUID generated by the system for realizing the entity object. In most cases this should be same as 'unique_id' of the entity. However, in some cases this can be different because of entities have migrated thier unique identifier to NSX Policy intent objects later in the timeline and did not use unique_id for realization. Realization id is helpful for users to debug data path to correlate the configuration with corresponding intent. 
	RealizationId string `json:"realization_id,omitempty"`
	// Path relative from its parent
	RelativePath string `json:"relative_path,omitempty"`
	// This is a UUID generated by the GM/LM to uniquely identify entites in a federated environment. For entities that are stretched across multiple sites, the same ID will be used on all the stretched sites. 
	UniqueId string `json:"unique_id,omitempty"`
	// subtree for this type within policy tree containing nested elements. 
	Children []ChildPolicyConfigResource `json:"children,omitempty"`
	// Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects. 
	MarkedForDelete bool `json:"marked_for_delete,omitempty"`
	// Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties. 
	Overridden bool `json:"overridden,omitempty"`
	// Flag to enable ECMP. 
	Ecmp bool `json:"ecmp,omitempty"`
	// Flag to enable BGP configuration. Disabling will stop feature and BGP peering. 
	Enabled bool `json:"enabled,omitempty"`
	// Flag to enable graceful restart. This field is deprecated, please use graceful_restart_config parameter for graceful restart configuration. If both parameters are set and consistent with each other (i.e. graceful_restart=false and graceful_restart_mode=HELPER_ONLY OR graceful_restart=true and graceful_restart_mode=GR_AND_HELPER) then this is allowed, but if inconsistent with each other then this is not allowed and validation error will be thrown. 
	GracefulRestart bool `json:"graceful_restart,omitempty"`
	GracefulRestartConfig *BgpGracefulRestartConfig `json:"graceful_restart_config,omitempty"`
	// Flag to enable inter SR IBGP configuration. When not specified, inter SR IBGP is automatically enabled if Tier-0 is created in ACTIVE_ACTIVE ha_mode. 
	InterSrIbgp bool `json:"inter_sr_ibgp,omitempty"`
	// Specify BGP AS number for Tier-0 to advertize to BGP peers. AS number can be specified in ASPLAIN (e.g., \"65546\") or ASDOT (e.g., \"1.10\") format. Empty string disables BGP feature. It is required by normal tier0 but not required in vrf tier0. 
	LocalAsNum string `json:"local_as_num,omitempty"`
	// Flag to enable BGP multipath relax option.
	MultipathRelax bool `json:"multipath_relax,omitempty"`
	// List of routes to be aggregated. 
	RouteAggregations []RouteAggregationEntry `json:"route_aggregations,omitempty"`
}
