/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// VIF attachment state of a segment port
type SegmentPortAttachmentState struct {
	// VM or vmknic entities that are attached to the Segment Port
	Attachers []PortAttacher `json:"attachers,omitempty"`
	// VIF ID
	Id string `json:"id,omitempty"`
	// A segment port must be in one of following states. FREE - If there are no active attachers. The port may or may not have an attachment ID configured on it. This state is applicable only to port of static type. ATTACHED - Segment port has exactly one active attacher and no further configuration is pending. ATTACHED_PENDING_CONF - Segment port has exactly one attacher, however it may not have been configured completely. Additional configuration will be provided by other nsx components. ATTACHED_IN_MOTION - Segment port has multiple active attachers. This state represents a scenario where VM is moving from one location (host or storage) to another (e.g. vmotion, vSphere HA) DETACHED - A temporary state after all port attachers have been detached. This state is applicable only to a port of ephemeral type and the port will soon be deleted. 
	State string `json:"state,omitempty"`
}
