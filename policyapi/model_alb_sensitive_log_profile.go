/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer SensitiveLogProfile object
type AlbSensitiveLogProfile struct {
	// Match sensitive header fields in HTTP application log.
	HeaderFieldRules []AlbSensitiveFieldRule `json:"header_field_rules,omitempty"`
	// Match sensitive URI query params in HTTP application log. Query params from the URI are extracted and checked for matching sensitive parameter names. A successful match will mask the parameter values in accordance with this rule action. 
	UriQueryFieldRules []AlbSensitiveFieldRule `json:"uri_query_field_rules,omitempty"`
	// Match sensitive WAF log fields in HTTP application log.
	WafFieldRules []AlbSensitiveFieldRule `json:"waf_field_rules,omitempty"`
}
