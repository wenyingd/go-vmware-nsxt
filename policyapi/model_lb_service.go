/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Loadbalancer Service.
type LbService struct {
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
	// Flag to enable access log
	AccessLogEnabled bool `json:"access_log_enabled,omitempty"`
	// LBS could be instantiated (or created) on the Tier-1, etc. For now, only the Tier-1 object is supported. 
	ConnectivityPath string `json:"connectivity_path,omitempty"`
	// Flag to enable the load balancer service.
	Enabled bool `json:"enabled,omitempty"`
	// Load balancer engine writes information about encountered issues of different severity levels to the error log. This setting is used to define the severity level of the error log. 
	ErrorLogLevel string `json:"error_log_level,omitempty"`
	// If relax_scale_validation is true, the scale validations for virtual servers/pools/pool members/rules are relaxed for load balancer service. When load balancer service is deployed on edge nodes, the scale of virtual servers/pools/pool members for the load balancer service should not exceed the scale number of the largest load balancer size which could be configured on a certain edge form factor. For example, the largest load balancer size supported on a MEDIUM edge node is MEDIUM. So one SMALL load balancer deployed on MEDIUM edge nodes can support the scale number of MEDIUM load balancer. It is not recommended to enable active monitors if relax_scale_validation is true due to performance consideration. If relax_scale_validation is false, scale numbers should be validated for load balancer service. The property is deprecated as NSX-T Load Balancer is deprecated. 
	RelaxScaleValidation bool `json:"relax_scale_validation,omitempty"`
	// Load balancer service size. The load balancer service sizes, SMALL, MEDIUM, LARGE and XLARGE are all deprecated. Customers who are using this set of features are advised to migrate to NSX Advanced Load Balancer (Avi) which provides a superset of the NSX-T load balancing functionality. 
	Size string `json:"size,omitempty"`
}
