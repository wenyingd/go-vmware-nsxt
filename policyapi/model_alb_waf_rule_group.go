/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer WafRuleGroup object
type AlbWafRuleGroup struct {
	// Enable or disable WAF Rule Group. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	Enable bool `json:"enable,omitempty"`
	// Exclude list for the WAF rule group. The fields in the exclude list entry are logically and'ed to deduce the exclusion criteria. If there are multiple excludelist entries, it will be 'logical or' of them. Maximum of 64 items allowed. 
	ExcludeList []AlbWafExcludeListEntry `json:"exclude_list,omitempty"`
	// Number of index.
	Index int64 `json:"index"`
	// Name of the object.
	Name string `json:"name"`
	// Rules as per Modsec language. Maximum of 1024 items allowed. 
	Rules []AlbWafRule `json:"rules,omitempty"`
}
