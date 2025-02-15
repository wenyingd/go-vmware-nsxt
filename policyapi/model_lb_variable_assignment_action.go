/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// This action is used to create a new variable and assign value to it. One action can be used to create one variable. To create multiple variables, multiple actions must be defined. The variables can be used by LBVariableCondition, etc. 
type LbVariableAssignmentAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Name of the variable to be assigned.
	VariableName string `json:"variable_name"`
	// Value of variable.
	VariableValue string `json:"variable_value"`
}
