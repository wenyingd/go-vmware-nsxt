/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer DnsRecord object
type AlbDnsRecord struct {
	// Specifies the algorithm to pick the IP address(es) to be returned, when multiple entries are configured. This does not apply if num_records_in_response is 0. Default is round-robin. Enum options - DNS_RECORD_RESPONSE_ROUND_ROBIN, DNS_RECORD_RESPONSE_CONSISTENT_HASH. Default value when not specified in API or module is interpreted by ALB Controller as DNS_RECORD_RESPONSE_ROUND_ROBIN. 
	Algorithm string `json:"algorithm,omitempty"`
	Cname *AlbDnsCnameRdata `json:"cname,omitempty"`
	// Configured FQDNs are delegated domains (i.e. they represent a zone cut). Default value when not specified in API or module is interpreted by ALB Controller as false. 
	Delegated bool `json:"delegated,omitempty"`
	// Details of DNS record.
	Description string `json:"description,omitempty"`
	// Fully Qualified Domain Name. Minimum of 1 items required. 
	Fqdn []string `json:"fqdn"`
	// IPv6 address in AAAA record. Maximum of 4 items allowed. 
	Ip6Address []AlbDnsAaaaRdata `json:"ip6_address,omitempty"`
	// IP address in A record. Maximum of 4 items allowed. 
	IpAddress []AlbDnsARdata `json:"ip_address,omitempty"`
	// Internal metadata for the DNS record.
	Metadata string `json:"metadata,omitempty"`
	// MX record. Maximum of 4 items allowed. 
	MxRecords []AlbDnsMxRdata `json:"mx_records,omitempty"`
	// Name Server information in NS record. Maximum of 13 items allowed. 
	Ns []AlbDnsNsRdata `json:"ns,omitempty"`
	// Specifies the number of records returned by the DNS service. Enter 0 to return all records. Default is 0. Allowed values are 0-20. Special values are 0- 'Return all records'. Default value when not specified in API or module is interpreted by ALB Controller as 0. 
	NumRecordsInResponse int64 `json:"num_records_in_response,omitempty"`
	// Service locator info in SRV record. Maximum of 4 items allowed. 
	ServiceLocator []AlbDnsSrvRdata `json:"service_locator,omitempty"`
	// Time To Live for this DNS record.
	Ttl int64 `json:"ttl,omitempty"`
	// Text record. Maximum of 4 items allowed. 
	TxtRecords []AlbDnsTxtRdata `json:"txt_records,omitempty"`
	// DNS record type. Enum options - DNS_RECORD_OTHER, DNS_RECORD_A, DNS_RECORD_NS, DNS_RECORD_CNAME, DNS_RECORD_SOA, DNS_RECORD_PTR, DNS_RECORD_HINFO, DNS_RECORD_MX, DNS_RECORD_TXT, DNS_RECORD_RP, DNS_RECORD_DNSKEY, DNS_RECORD_AAAA, DNS_RECORD_SRV, DNS_RECORD_OPT, DNS_RECORD_RRSIG, DNS_RECORD_AXFR, DNS_RECORD_ANY. 
	Type_ string `json:"type"`
	// Enable wild-card match of fqdn  if an exact match is not found in the DNS table, the longest match is chosen by wild-carding the fqdn in the DNS request. Default is false. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	WildcardMatch bool `json:"wildcard_match,omitempty"`
}