/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Action to be perform on deployment.
type CloudNativeDeploymentAction struct {
	// Action can be deploy or undeploy. DEPLOY - Deploy NSX Application Platform charts. UNDEPLOY - Undeploy NSX Application Platform charts. REDEPLOY - Redeploy NSX Application Platform charts. UPDATE_FORMFACTOR - Upgrade NSX Application Platform charts. REDEPLOY_UPDATE_FORMFACTOR - Retry update NSX Application Platform charts. FORCE_UNDEPLOY - Undeploy forcefully. RESTART - Restart deployment. RESET - Reset deployment. 
	Action string `json:"action"`
}