/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type IntervalSampling struct {
	// Sampling type
	SamplingType string `json:"sampling_type"`
	// Time interval in ms between two sampling actions.
	SamplingInterval int64 `json:"sampling_interval"`
}
