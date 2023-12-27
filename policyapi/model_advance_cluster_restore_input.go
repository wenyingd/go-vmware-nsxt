/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type AdvanceClusterRestoreInput struct {
	// Unique id of an instruction (as returned by the GET /restore/status call) for which input is to be provided 
	Id string `json:"id,omitempty"`
	// List of resources for which the instruction is applicable.
	Resources []SelectableResourceReference `json:"resources"`
}