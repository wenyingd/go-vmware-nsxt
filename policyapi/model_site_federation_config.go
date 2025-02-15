/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Site fedeation configuration.
type SiteFederationConfig struct {
	// Remote tunnel endpoint IP addresses
	RtepIps []string `json:"rtep_ips,omitempty"`
	// Site UUID
	SiteId string `json:"site_id,omitempty"`
	// Unique site index allocated (from range 0-4095)
	SiteIndex int64 `json:"site_index,omitempty"`
	// Site path
	SitePath string `json:"site_path,omitempty"`
}
