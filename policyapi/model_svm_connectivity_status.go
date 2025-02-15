/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type SvmConnectivityStatus struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	Self *SelfResourceLink `json:"_self,omitempty"`
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
	Tags []Tag `json:"tags,omitempty"`
	// Connectivity status with the deployed Solution VM TRUE  - VM is configured and protected by EPP/AMS Service VM. FALSE - VM is either not configured for protection or VM is disconnected from EPP/AMS Service VM.
	ConnectivityStatus bool `json:"connectivity_status,omitempty"`
	// Service name as provided for Anti Malware Solution or as provided for third party Endpoint Protection solution during service registration.
	ServiceName string `json:"service_name,omitempty"`
	// Solution ID as provided for Anti Malware Solution(AMS) or as provided for third party Endpoint Protection(EPP) solution during service registration.
	SolutionId string `json:"solution_id,omitempty"`
}
