/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer ConnPoolProperties object
type AlbConnPoolProperties struct {
	// Connection idle timeout. Allowed in Basic(Allowed values- 60000) edition, Essentials(Allowed values- 60000) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as 60000. 
	UpstreamConnpoolConnIdleTmo int64 `json:"upstream_connpool_conn_idle_tmo,omitempty"`
	// Connection life timeout. Allowed in Basic(Allowed values- 600000) edition, Essentials(Allowed values- 600000) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as 600000. 
	UpstreamConnpoolConnLifeTmo int64 `json:"upstream_connpool_conn_life_tmo,omitempty"`
	// Maximum number of times a connection can be reused. Special values are 0- 'unlimited'. Allowed in Basic(Allowed values- 0) edition, Essentials(Allowed values- 0) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as 0. 
	UpstreamConnpoolConnMaxReuse int64 `json:"upstream_connpool_conn_max_reuse,omitempty"`
	// Maximum number of connections a server can cache. Special values are 0- 'unlimited'. Default value when not specified in API or module is interpreted by ALB Controller as 0. 
	UpstreamConnpoolServerMaxCache int64 `json:"upstream_connpool_server_max_cache,omitempty"`
}
