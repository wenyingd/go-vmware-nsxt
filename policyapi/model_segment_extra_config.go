/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Segment extra config is intended for supporting vendor specific configuration on the data path, it can be set as key value string pairs on either segment or segment port. 
type SegmentExtraConfig struct {
	ConfigPair *UnboundedKeyValuePair `json:"config_pair"`
}
