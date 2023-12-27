/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Dropdown item definition
type DropdownItem struct {
	// An additional key-value pair for item to be display in dropdown.
	AdditionalValue interface{} `json:"additional_value,omitempty"`
	// expression to extract display name to be shown in the drop down.
	DisplayName string `json:"display_name,omitempty"`
	// An expression that represents the items of the dropdown filter.
	Field string `json:"field"`
	// Property value is shown in the drop down input box for a filter. If the value is not provided 'display_name' property value is used.
	ShortDisplayName string `json:"short_display_name,omitempty"`
	// Value of filter inside dropdown filter.
	Value string `json:"value"`
}