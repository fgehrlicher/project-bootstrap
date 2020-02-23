package parser

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"project-bootstrap/types"
)

// Parser parses definitions
type Parser struct{}

// GetKind returns the kind of a file handle
func (parser Parser) GetKind(handle io.Reader) (string, error) {
	baseDefinition := types.BaseDefinition{}
	content, err := ioutil.ReadAll(handle)
	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal(content, &baseDefinition)
	if err != nil {
		return "", err
	}

	kind := baseDefinition.Kind
	if !types.AllowedKinds.IsAllowed(kind) {
		return "", types.ErrInvalidKind
	}

	return kind, err
}

// ParseWorkflow parses a workflow
func (parser Parser) ParseWorkflow(handle io.Reader) (types.Workflow, error) {
	workflow := types.Workflow{BaseDefinition: types.BaseDefinition{}}
	content, err := ioutil.ReadAll(handle)
	if err != nil {
		return workflow, err
	}

	err = yaml.Unmarshal(content, &workflow)
	return workflow, err
}

// ParseAction parses a action
func (parser Parser) ParseAction(handle io.Reader) (types.Action, error) {
	action := types.Action{BaseDefinition: types.BaseDefinition{}}
	content, err := ioutil.ReadAll(handle)
	if err != nil {
		return action, err
	}

	err = yaml.Unmarshal(content, &action)
	return action, err
}
