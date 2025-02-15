/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HSMAwsCloudHsm object
type AlbhsmAwsCloudHsm struct {
	// client_config of HSMAwsCloudHsm.
	ClientConfig string `json:"client_config,omitempty"`
	// AWS CloudHSM Cluster Certificate.
	ClusterCert string `json:"cluster_cert,omitempty"`
	// Username of the Crypto User. This will be used to access the keys on the HSM . 
	CryptoUserName string `json:"crypto_user_name,omitempty"`
	// Password of the Crypto User. This will be used to access the keys on the HSM . 
	CryptoUserPassword string `json:"crypto_user_password,omitempty"`
	// IP address of the HSM in the cluster. If there are more than one HSMs, only one is sufficient. 
	HsmIp []string `json:"hsm_ip,omitempty"`
	// mgmt_config of HSMAwsCloudHsm.
	MgmtConfig string `json:"mgmt_config,omitempty"`
}
