/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer WafRuleGroupOverrides object
type AlbWafRuleGroupOverrides struct {
	// Override the enable flag for this group.
	Enable bool `json:"enable,omitempty"`
	// Replace the exclude list for this group. Maximum of 64 items allowed. 
	ExcludeList []AlbWafExcludeListEntry `json:"exclude_list,omitempty"`
	// Override the waf mode for this group. Enum options - WAF_MODE_DETECTION_ONLY, WAF_MODE_ENFORCEMENT. 
	Mode string `json:"mode,omitempty"`
	// The name of the group where attributes or rules are overridden. 
	Name string `json:"name"`
	// Rule specific overrides. Maximum of 1024 items allowed. 
	RuleOverrides []AlbWafRuleOverrides `json:"rule_overrides,omitempty"`
}