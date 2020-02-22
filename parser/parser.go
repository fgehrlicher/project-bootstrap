package parser

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"project-bootstrap/types"
)

// Parser parses stuff
type Parser struct{}

// ParseWorkflow parses a workflow
func (parser Parser) ParseWorkflow(handle io.Reader) (workflow types.Workflow, err error) {
	workflow = types.Workflow{BaseDefinition: types.BaseDefinition{}}
	content, err := ioutil.ReadAll(handle)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(content, &workflow)
	return
}
