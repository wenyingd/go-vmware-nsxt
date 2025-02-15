/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Status LDAP server of directory domain
type DirectoryLdapServerStatus struct {
	// Error ID of the directory LDAP server status maintained by the NSX directory service.
	ErrorId int64 `json:"error_id,omitempty"`
	// Error message of the directory LDAP server status maintained by the NSX directory service.
	ErrorMessage string `json:"error_message,omitempty"`
}
