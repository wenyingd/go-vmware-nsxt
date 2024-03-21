/*
 * NSX Policy API
 *
 * VMware NSX Policy REST API
 *
 * API version: 4.2.0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Paged collection of SegmentConnectionBindingMaps
type SegmentConnectionBindingMapListResult struct {
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	// Opaque cursor to be used for getting next page of records (supplied by current result page)
	Cursor string `json:"cursor,omitempty"`
	// If true, results are sorted in ascending order
	SortAscending bool `json:"sort_ascending,omitempty"`
	// Field by which records are sorted
	SortBy string `json:"sort_by,omitempty"`
	// Count of results found (across all pages), set only on first page
	ResultCount int64 `json:"result_count,omitempty"`
	// Segment connection binding map list results
	Results []SegmentConnectionBindingMap `json:"results"`
}