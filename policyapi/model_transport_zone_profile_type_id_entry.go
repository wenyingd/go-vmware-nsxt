/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type TransportZoneProfileTypeIdEntry struct {
	// profile id of the resource type
	ProfileId string `json:"profile_id"`
	// Selects the type of the transport zone profile
	ResourceType string `json:"resource_type,omitempty"`
}
