package types

// ActionType is the string which defines an action
const ActionType = "action"

// Action is the abstract action definition
type Action struct {
	BaseDefinition `yaml:",inline"`
}
