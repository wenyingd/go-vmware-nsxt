/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// L2VPNSessionPeerCodes represents an array of peer code for each tunnel. The peer code is necessary to configure the remote end of the tunnel. Currently only stand-along/unmanaged edge is supported on the remote end of the tunnel. 
type L2VpnSessionPeerConfigNsxT struct {
	Alarm *PolicyRuntimeAlarm `json:"alarm,omitempty"`
	// Policy Path referencing the enforcement point where the info is fetched. 
	EnforcementPointPath string `json:"enforcement_point_path,omitempty"`
	ResourceType string `json:"resource_type"`
	// List of peer codes per transport tunnel.
	PeerCodes []L2VpnSessionTransportTunnelPeerCode `json:"peer_codes,omitempty"`
}
