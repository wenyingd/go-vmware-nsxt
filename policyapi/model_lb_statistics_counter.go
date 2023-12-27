/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type LbStatisticsCounter struct {
	// Number of bytes in.
	BytesIn int64 `json:"bytes_in,omitempty"`
	// The average number of inbound bytes per second, the number is averaged over the last 5 one-second intervals. 
	BytesInRate float32 `json:"bytes_in_rate,omitempty"`
	// Number of bytes out.
	BytesOut int64 `json:"bytes_out,omitempty"`
	// The average number of outbound bytes per second, the number is averaged over the last 5 one-second intervals. 
	BytesOutRate float32 `json:"bytes_out_rate,omitempty"`
	// The average number of current sessions per second, the number is averaged over the last 5 one-second intervals. 
	CurrentSessionRate float32 `json:"current_session_rate,omitempty"`
	// Number of current sessions.
	CurrentSessions int64 `json:"current_sessions,omitempty"`
	// The total number of dropped TCP SYN or UDP packets by access list control. 
	DroppedPacketsByAccessList int64 `json:"dropped_packets_by_access_list,omitempty"`
	// The total number of dropped sessions by LB rule action. 
	DroppedSessionsByLbruleAction int64 `json:"dropped_sessions_by_lbrule_action,omitempty"`
	// The average number of http requests per second, the number is averaged over the last 5 one-second intervals. 
	HttpRequestRate float32 `json:"http_request_rate,omitempty"`
	// The total number of http requests.
	HttpRequests int64 `json:"http_requests,omitempty"`
	// Number of maximum sessions.
	MaxSessions int64 `json:"max_sessions,omitempty"`
	// Number of packets in.
	PacketsIn int64 `json:"packets_in,omitempty"`
	// The average number of inbound packets per second, the number is averaged over the last 5 one-second intervals. 
	PacketsInRate float32 `json:"packets_in_rate,omitempty"`
	// Number of packets out.
	PacketsOut int64 `json:"packets_out,omitempty"`
	// The average number of outbound packets per second, the number is averaged over the last 5 one-second intervals. 
	PacketsOutRate float32 `json:"packets_out_rate,omitempty"`
	// Number of source IP persistence entries
	SourceIpPersistenceEntrySize int64 `json:"source_ip_persistence_entry_size,omitempty"`
	// Number of total sessions.
	TotalSessions int64 `json:"total_sessions,omitempty"`
}