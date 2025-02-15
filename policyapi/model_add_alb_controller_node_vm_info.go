/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Contains a list of Advanced Load Balancer controller node VM deployment requests. 
type AddAlbControllerNodeVmInfo struct {
	// Advanced Load Balancer Controller deployment requests to be deployed by NSX. 
	DeploymentRequests []AlbControllerNodeVmDeploymentRequest `json:"deployment_requests"`
}
