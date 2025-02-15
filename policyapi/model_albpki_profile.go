/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer PKIProfile object
type AlbpkiProfile struct {
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
	// List of Certificate Authorities (Root and Intermediate) trusted that is used for certificate validation. 
	CaCerts []AlbsslCertificate `json:"ca_certs,omitempty"`
	// Creator name.
	CreatedBy string `json:"created_by,omitempty"`
	// When enabled, Avi will verify via CRL checks that certificates in the trust chain have not been revoked. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	CrlCheck bool `json:"crl_check,omitempty"`
	// Certificate Revocation Lists.
	Crls []Albcrl `json:"crls,omitempty"`
	// When enabled, Avi will not trust Intermediate and Root certs presented by a client. Instead, only the chain certs configured in the Certificate Authority section will be used to verify trust of the client's cert. Allowed in Basic(Allowed values- true) edition, Essentials(Allowed values- true) edition, Enterprise edition. Special default for Basic edition is true, Essentials edition is true, Enterprise is False. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	IgnorePeerChain bool `json:"ignore_peer_chain,omitempty"`
	// This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	IsFederated bool `json:"is_federated,omitempty"`
	// List of labels to be used for granular RBAC. Allowed in Basic edition, Essentials edition, Enterprise edition. 
	Markers []AlbRoleFilterMatchLabel `json:"markers,omitempty"`
	// When enabled, Avi will only validate the revocation status of the leaf certificate using CRL. To enable validation for the entire chain, disable this option and provide all the relevant CRLs. Allowed in Basic(Allowed values- true) edition, Essentials(Allowed values- true) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	ValidateOnlyLeafCrl bool `json:"validate_only_leaf_crl,omitempty"`
}
