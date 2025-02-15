/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Combines two or more widgetconfigurations into a multi-widget
type MultiWidgetConfiguration struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Title of the widget. If display_name is omitted, the widget will be shown without a title.
	DisplayName string `json:"display_name,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// Supported visualization types are LabelValueConfiguration, DonutConfiguration, GridConfiguration, StatsConfiguration, MultiWidgetConfiguration, GraphConfiguration, ContainerConfiguration, CustomWidgetConfiguration, CustomFilterWidgetConfiguration, TimeRangeDropdownFilterWidgetConfiguration, SpacerWidgetConfiguration, LegendWidgetConfiguration and DropdownFilterWidgetConfiguration.
	ResourceType string `json:"resource_type"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// If the condition is met then the widget will be displayed to UI. If no condition is provided, then the widget will be displayed unconditionally.
	Condition string `json:"condition,omitempty"`
	// The 'datasources' represent the sources from which data will be fetched. Currently, only NSX-API is supported as a 'default' datasource. An example of specifying 'default' datasource along with the urls to fetch data from is given at 'example_request' section of 'CreateWidgetConfiguration' API.
	Datasources []Datasource `json:"datasources,omitempty"`
	// Default filter values to be passed to datasources. This will be used when the report is requested without filter values.
	DefaultFilterValue []DefaultFilterValue `json:"default_filter_value,omitempty"`
	// Id of drilldown widget, if any. Id should be a valid id of an existing widget. A widget is considered as drilldown widget when it is associated with any other widget and provides more detailed information about any data item from the parent widget.
	DrilldownId string `json:"drilldown_id,omitempty"`
	FeatureSet *FeatureSet `json:"feature_set,omitempty"`
	// Id of filter widget for subscription, if any. Id should be a valid id of an existing filter widget. Filter widget should be from the same view. Datasource URLs should have placeholder values equal to filter alias to accept the filter value on filter change. This field is deprecated instead use 'filters' property.
	Filter string `json:"filter,omitempty"`
	// Flag to indicate that widget will continue to work without filter value. If this flag is set to false then default_filter_value is manadatory.
	FilterValueRequired bool `json:"filter_value_required,omitempty"`
	// A List of filter applied to this widget configuration. This will be used to identify the filters applied to this widget.
	Filters []string `json:"filters,omitempty"`
	Footer *Footer `json:"footer,omitempty"`
	// Icons to be applied at dashboard for widgets and UI elements.
	Icons []Icon `json:"icons,omitempty"`
	// Set to true if this widget should be used as a drilldown.
	IsDrilldown bool `json:"is_drilldown,omitempty"`
	Legend *Legend `json:"legend,omitempty"`
	// List of plotting configuration for a given widget. Widget plotting configurations which are common across all the widgets types should be define here.
	PlotConfigs []WidgetPlotConfiguration `json:"plot_configs,omitempty"`
	// Represents the vertical span of the widget / container. 1 Row span is equal to 20px.
	Rowspan int32 `json:"rowspan,omitempty"`
	// Please use the property 'shared' of View instead of this. The widgets of a shared view are visible to other users.
	Shared bool `json:"shared,omitempty"`
	// If the value of this field is set to true then card header will be displayed otherwise only card will be displayed without header.
	ShowHeader bool `json:"show_header,omitempty"`
	// Represents the horizontal span of the widget / container.
	Span int32 `json:"span,omitempty"`
	// Specify relavite weight in WidgetItem for placement in a view. Please see WidgetItem for details.
	Weight int32 `json:"weight,omitempty"`
	// Hyperlink of the specified UI page that provides details.
	Navigation string `json:"navigation,omitempty"`
	// Array of widgets that are part of the multi-widget.
	Widgets []WidgetItem `json:"widgets"`
}
