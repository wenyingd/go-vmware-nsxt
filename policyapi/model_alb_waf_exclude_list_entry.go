/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer WafExcludeListEntry object
type AlbWafExcludeListEntry struct {
	ClientSubnet *AlbIpAddrPrefix `json:"client_subnet,omitempty"`
	// Free-text comment about this exclusion.
	Description string `json:"description,omitempty"`
	// The match_element can be 'ARGS xxx', 'ARGS_GET xxx', 'ARGS_POST xxx', 'ARGS_NAMES xxx', 'FILES xxx', 'QUERY_STRING', 'REQUEST_BASENAME', 'REQUEST_BODY', 'REQUEST_URI', 'REQUEST_URI_RAW', 'REQUEST_COOKIES xxx', 'REQUEST_HEADERS xxx' or 'RESPONSE_HEADERS xxx'. These match_elements in the HTTP Transaction (if present) will be excluded when executing WAF Rules. 
	MatchElement string `json:"match_element,omitempty"`
	MatchElementCriteria *AlbWafExclusionType `json:"match_element_criteria,omitempty"`
	UriMatchCriteria *AlbWafExclusionType `json:"uri_match_criteria,omitempty"`
	// URI Path to exclude for WAF rules.
	UriPath string `json:"uri_path,omitempty"`
}