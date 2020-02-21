package types

// WorkflowType is the string which defines a workflow
const WorkflowType = "workflow"

// Workflow is the abstract workflow definition
type Workflow struct {
	*BaseDefinition
}
