/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type RaDnsConfig struct {
	// DNS server. 
	DnsServer []string `json:"dns_server,omitempty"`
	// Lifetime of DNS server in milliseconds
	DnsServerLifetime int64 `json:"dns_server_lifetime,omitempty"`
	// Domain name in RA message. 
	DomainName []string `json:"domain_name,omitempty"`
	// Lifetime of Domain names in milliseconds
	DomainNameLifetime int64 `json:"domain_name_lifetime,omitempty"`
}
