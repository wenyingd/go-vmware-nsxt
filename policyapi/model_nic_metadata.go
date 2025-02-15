/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Information on the Network interfaces present on the partner appliance that needs to be configured by the NSX Manager.
type NicMetadata struct {
	// Network Interface index.
	InterfaceIndex int64 `json:"interface_index"`
	// Network Interface label.
	InterfaceLabel string `json:"interface_label"`
	// Interface that needs to be configured on the partner appliance. Ex. MANAGEMENT, DATA1, DATA2, HA1, HA2, CONTROL.
	InterfaceType string `json:"interface_type"`
	// Transport Type of the service, which is the mechanism of redirecting the traffic to the the partner appliance. Transport type is required if Service caters to any functionality other than EPP and MPS. Here, the transports array specifies the kinds of transport where this particular NIC is user configurable. If nothing is specified, and the \"user_configurable\" flag is true, then user configuration will be allowed for all transports. If any transport is/are specified, then it will be considered as user configurable for the specified transports only.\"
	Transports []string `json:"transports,omitempty"`
	// Used to specify if the given interface needs configuration. Management nics will always need the configuration, for others it will be use case specific. For example, a DATA NIC may be user configurable if the appliance is deployed in certain mode, such as L3_ROUTED.
	UserConfigurable bool `json:"user_configurable,omitempty"`
}
