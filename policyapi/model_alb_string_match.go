/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer StringMatch object
type AlbStringMatch struct {
	// Criterion to use for string matching the HTTP request. Enum options - BEGINS_WITH, DOES_NOT_BEGIN_WITH, CONTAINS, DOES_NOT_CONTAIN, ENDS_WITH, DOES_NOT_END_WITH, EQUALS, DOES_NOT_EQUAL, REGEX_MATCH, REGEX_DOES_NOT_MATCH. Allowed in Basic(Allowed values- BEGINS_WITH,DOES_NOT_BEGIN_WITH,CONTAINS,DOES_NOT_CONTAIN,ENDS_WITH,DOES_NOT_END_WITH,EQUALS,DOES_NOT_EQUAL) edition, Essentials(Allowed values- BEGINS_WITH,DOES_NOT_BEGIN_WITH,CONTAINS,DOES_NOT_CONTAIN,ENDS_WITH,DOES_NOT_END_WITH,EQUALS,DOES_NOT_EQUAL) edition, Enterprise edition. 
	MatchCriteria string `json:"match_criteria"`
	// String value(s).
	MatchStr []string `json:"match_str,omitempty"`
	// path of the string group(s). It is a reference to an object of type StringGroup. 
	StringGroupPaths []string `json:"string_group_paths,omitempty"`
}
