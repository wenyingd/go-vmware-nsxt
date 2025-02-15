/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// If Virtual Distributed Switch is used as a HostSwitch to configure TransportNode or TransportNodeProfie, this mapping should be specified. You can either use vds_uplink_name or vds_lag_name to associate with uplink_name from UplinkHostSwitch profile.
type VdsUplink struct {
	// This name is from UplinkHostSwitch profile that is associated with the HostSwitch specified in TransportNode or TransportNodeProfile configuration. This name will be used as an alias to either VDS uplink or lag in other configuration.
	UplinkName string `json:"uplink_name"`
	// LAG name that is connected to Physical NIC on a host from vSphere.
	VdsLagName string `json:"vds_lag_name,omitempty"`
	// Uplink name of VDS that is connected to Physical NIC on a host from vSphere.
	VdsUplinkName string `json:"vds_uplink_name,omitempty"`
}
