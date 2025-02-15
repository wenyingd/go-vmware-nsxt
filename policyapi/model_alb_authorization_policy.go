/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer AuthorizationPolicy object
type AlbAuthorizationPolicy struct {
	// Authorization Policy Rules.
	AuthzRules []AlbAuthorizationRule `json:"authz_rules,omitempty"`
}
