/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HTTPRewriteLocHdrAction object
type AlbhttpRewriteLocHdrAction struct {
	Host *AlburiParam `json:"host,omitempty"`
	// Keep or drop the query from the server side redirect URI. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	KeepQuery bool `json:"keep_query,omitempty"`
	Path *AlburiParam `json:"path,omitempty"`
	// Port to use in the redirected URI. Allowed values are 1-65535. 
	Port int64 `json:"port,omitempty"`
	// HTTP protocol type. Enum options - HTTP, HTTPS. 
	Protocol string `json:"protocol"`
}