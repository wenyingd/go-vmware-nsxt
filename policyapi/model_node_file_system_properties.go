/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// File system properties
type NodeFileSystemProperties struct {
	// File system id
	FileSystem string `json:"file_system,omitempty"`
	// File system mount
	Mount string `json:"mount,omitempty"`
	// File system size in kilobytes
	Total int64 `json:"total,omitempty"`
	// File system type
	Type_ string `json:"type,omitempty"`
	// Amount of file system used in kilobytes
	Used int64 `json:"used,omitempty"`
}
