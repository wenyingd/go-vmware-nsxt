/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Represents X and Y axes of a graph. For a multi-graph, the same axes are shared by all the graphs.
type Axes struct {
	XLabel *Label `json:"x_label,omitempty"`
	// A list of X-Axis Labels with condition support. If needed, this property can be used to provide a list of x-axis label with condition support. For a label with single condition,'x-label' property can be used.
	XLabels []Label `json:"x_labels,omitempty"`
	// A list of Y-Axis unit Labels with condition support. If needed, this property can be used to provide a list of y-axis unit label with condition support. This unit label can be used to display the point value along with units like percentage, milliseconds etc.
	YAxisUnitLabels []Label `json:"y_axis_unit_labels,omitempty"`
	// A list of Y-Axis unit with condition support. If needed, this property can be used to provide a list of y-axis unit with condition support. This unit could be like percentage, seconds, milliseconds etc.
	YAxisUnits []AxisUnit `json:"y_axis_units,omitempty"`
	YLabel *Label `json:"y_label,omitempty"`
	// A list of Y-Axis Labels with condition support. If needed, this property can be used to provide a list of y-axis label with condition support. For a label with single condition,'y-label' property can be used.
	YLabels []Label `json:"y_labels,omitempty"`
}