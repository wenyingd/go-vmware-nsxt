/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Bfd Health Monitoring Options used specific to BFD Transport Zone profiles 
type BfdHealthMonitoringConfig struct {
	// Whether the heartbeat is enabled. A PATCH or PUT request with \"enabled\" false (with no probe intervals) will set or reset the probe_interval to their default value.
	Enabled bool `json:"enabled"`
	// The flag is to turn on/off latency. A PATCH or PUT request with \"latency_enabled\" true will enable NSX to send the networking latency data to thrid-party monitoring tools like vRNI.
	LatencyEnabled bool `json:"latency_enabled,omitempty"`
	// The time interval (in millisec) between probe packets for tunnels between transport nodes.
	ProbeInterval int64 `json:"probe_interval,omitempty"`
}
