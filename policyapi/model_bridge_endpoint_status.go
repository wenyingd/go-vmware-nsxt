/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type BridgeEndpointStatus struct {
	// The Ids of the transport nodes which actively serve the endpoint.
	ActiveNodes []string `json:"active_nodes,omitempty"`
	// The id of the bridge endpoint
	EndpointId string `json:"endpoint_id,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
}
