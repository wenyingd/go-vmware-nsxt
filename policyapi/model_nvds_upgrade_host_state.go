/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Individual host upgrade state
type NvdsUpgradeHostState struct {
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
	// DiscoveredNode identifier
	DnExtId string `json:"dn_ext_id,omitempty"`
	// TransportNode identifier
	Host string `json:"host,omitempty"`
	// TransportNode ip address
	IpAddress string `json:"ip_address,omitempty"`
	// Overall state of N-VDSes on the TransportNodes
	OverallState string `json:"overall_state,omitempty"`
	// Details of the N-VDS upgrade state on the host
	StateDetails []string `json:"state_details,omitempty"`
	// This field returns current stage of Migration task. Here is a sequence of stages the task cycles through, TN_MIGRATION_TASK_IN_QUEUE RETRIEVE_SAVED_CONFIG, TN_VALIDATE, VMS_RETRIVAL, VMS_UNREGISTRATION, TN_STATELESS_WAIT_FOR_HP, DETACH_TNP, TNP_WAIT, TN_SEND_HS_MIGRATION_MSG, TN_ADD_HOST_TO_VDS, TN_UPDATE, TN_UPDATE_WAIT, TN_DELETE, TN_DELETE_WAIT, FN_DELETE_WAIT, TN_RECONFIG_HOST, TN_CREATE, TN_CREATE_WAIT, UPDATE_TNP_AND_APPLY, TN_EXIT_MM, VMS_REGISTRATION, VMS_REGISTRATION_WAIT, TN_MIGRATION_COMPLETED Depending on the type of host (stateful, stateless, Sddc, etc.) migration task may not cycle through all stages but in will follow above sequence. If stage is TN_MIGRATION_COMPLETED refer to field overall_state for SUCCESS or UPGRADE_FAILURE and state_details for details on same. 
	UpgradeStage string `json:"upgrade_stage,omitempty"`
}
