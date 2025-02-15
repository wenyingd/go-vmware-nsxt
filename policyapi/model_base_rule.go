/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// A rule indicates the action to be performed for various types of traffic flowing between workload groups.
type BaseRule struct {
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
	// We need paths as duplicate names may exist for groups under different domains. Along with paths we support IP Address of type IPv4 and IPv6. IP Address can be in one of the format(CIDR, IP Address, Range of IP Address). In order to specify all groups, use the constant \"ANY\". This is case insensitive. If \"ANY\" is used, it should be the ONLY element in the group array. Error will be thrown if ANY is used in conjunction with other values. 
	DestinationGroups []string `json:"destination_groups,omitempty"`
	// If set to true, the rule gets applied on all the groups that are NOT part of the destination groups. If false, the rule applies to the destination groups 
	DestinationsExcluded bool `json:"destinations_excluded,omitempty"`
	// Define direction of traffic. 
	Direction string `json:"direction,omitempty"`
	// Flag to disable the rule. Default is enabled.
	Disabled bool `json:"disabled,omitempty"`
	// Type of IP packet that should be matched while enforcing the rule. The value is set to IPV4_IPV6 for Layer3 rule if not specified. For Layer2/Ether rule the value must be null. 
	IpProtocol string `json:"ip_protocol,omitempty"`
	// A flag to indicate whether rule is a default rule.
	IsDefault bool `json:"is_default,omitempty"`
	// Flag to enable packet logging. Default is disabled.
	Logged bool `json:"logged,omitempty"`
	// Text for additional notes on changes.
	Notes string `json:"notes,omitempty"`
	// Holds the list of layer 7 service profile paths. These profiles accept attributes and sub-attributes of various network services (e.g. L4 AppId, encryption algorithm, domain name, etc) as key value pairs. Instead of Layer 7 service profiles you can use a L7 access profile. One of either Layer 7 service profiles or L7 Access Profile can be used in firewall rule. In case of L7 access profile only one is allowed. 
	Profiles []string `json:"profiles,omitempty"`
	// This is a unique 4 byte positive number that is assigned by the system.  This rule id is passed all the way down to the data path. The first 1GB (1000 to 2^30) will be shared by GM and LM with zebra style striped number space. For E.g 1000 to (1Million -1) by LM, (1M - 2M-1) by GM and so on. 
	RuleId int64 `json:"rule_id,omitempty"`
	// The list of policy paths where the rule is applied LR/Edge/T0/T1/LRP etc. Note that a given rule can be applied on multiple LRs/LRPs. 
	Scope []string `json:"scope,omitempty"`
	// This field is used to resolve conflicts between multiple Rules under Security or Gateway Policy for a Domain If no sequence number is specified in the payload, a value of 0 is assigned by default. If there are multiple rules with the same sequence number then their order is not deterministic. If a specific order of rules is desired, then one has to specify unique sequence numbers or use the POST request on the rule entity with a query parameter action=revise to let the framework assign a sequence number 
	SequenceNumber int32 `json:"sequence_number,omitempty"`
	// In order to specify raw services this can be used, along with services which contains path to services. This can be empty or null. 
	ServiceEntries []ServiceEntry `json:"service_entries,omitempty"`
	// In order to specify all services, use the constant \"ANY\". This is case insensitive. If \"ANY\" is used, it should be the ONLY element in the services array. Error will be thrown if ANY is used in conjunction with other values. 
	Services []string `json:"services,omitempty"`
	// We need paths as duplicate names may exist for groups under different domains. Along with paths we support IP Address of type IPv4 and IPv6. IP Address can be in one of the format(CIDR, IP Address, Range of IP Address). In order to specify all groups, use the constant \"ANY\". This is case insensitive. If \"ANY\" is used, it should be the ONLY element in the group array. Error will be thrown if ANY is used in conjunction with other values. 
	SourceGroups []string `json:"source_groups,omitempty"`
	// If set to true, the rule gets applied on all the groups that are NOT part of the source groups. If false, the rule applies to the source groups 
	SourcesExcluded bool `json:"sources_excluded,omitempty"`
	// User level field which will be printed in CLI and packet logs. Even though there is no limitation on length of a tag, internally tag will get truncated after 32 characters. 
	Tag string `json:"tag,omitempty"`
}
