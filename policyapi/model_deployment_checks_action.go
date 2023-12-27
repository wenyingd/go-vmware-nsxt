/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Pre/Post deployment check.
type DeploymentChecksAction struct {
	// Run pre/post deployment checks. PRE_CHECKS - Run pre-check before deployment. POST_CHECKS - Run post-check after deployment. ABORT_CHECKS - Abort running pre/post deployement checks. 
	Action string `json:"action"`
}