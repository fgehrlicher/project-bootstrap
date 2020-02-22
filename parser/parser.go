package parser

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"project-bootstrap/types"
)

// Parser parses definitions
type Parser struct{}

// ParseWorkflow parses a workflow
func (parser Parser) ParseWorkflow(handle io.Reader) (workflow types.Workflow, err error) {
	workflow.BaseDefinition = types.BaseDefinition{}
	content, err := ioutil.ReadAll(handle)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(content, &workflow)
	return
}

// ParseWorkflow parses a action
func (parser Parser) ParseAction(handle io.Reader) (action types.Action, err error) {
	action.BaseDefinition = types.BaseDefinition{}
	content, err := ioutil.ReadAll(handle)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(content, &action)
	return
}
