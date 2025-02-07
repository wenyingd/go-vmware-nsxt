/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// An instance of a datasource configuration.
type Datasource struct {
	// Name of a datasource instance.
	DisplayName string `json:"display_name"`
	KeystoreInfo *KeyStoreInfo `json:"keystore_info,omitempty"`
	// Array of urls relative to the datasource configuration. For example, api/v1/fabric/nodes is a relative url of nsx-manager instance.
	Urls []UrlAlias `json:"urls"`
}
