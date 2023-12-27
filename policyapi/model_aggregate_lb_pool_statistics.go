/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Paged Collection of LBPoolStatisticsPerEP
type AggregateLbPoolStatistics struct {
	// Intent path of object, forward slashes must be escaped using %2F. 
	IntentPath string `json:"intent_path,omitempty"`
	// LBPoolStatisticsPerEP list results.
	Results []LbPoolStatisticsPerEp `json:"results,omitempty"`
}