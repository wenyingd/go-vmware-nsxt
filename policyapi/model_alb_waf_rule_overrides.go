/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer WafRuleOverrides object
type AlbWafRuleOverrides struct {
	// Override the enable flag for this rule.
	Enable bool `json:"enable,omitempty"`
	// Replace the exclude list for this rule. Maximum of 64 items allowed. 
	ExcludeList []AlbWafExcludeListEntry `json:"exclude_list,omitempty"`
	// Override the waf mode for this rule. Enum options - WAF_MODE_DETECTION_ONLY, WAF_MODE_ENFORCEMENT. 
	Mode string `json:"mode,omitempty"`
	// The rule_id of the rule where attributes are overridden.
	RuleId string `json:"rule_id"`
}