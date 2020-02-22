package parser

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"project-bootstrap/types"
)

type InvalidReader struct{}

var errInvalidReader = errors.New("im an invalid reader")

func (i InvalidReader) Read(p []byte) (n int, err error) {
	err = errInvalidReader
	return
}

func loadFixture(t *testing.T, name string) *os.File {
	file, err := os.Open("./testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}

	return file
}

func TestParserParseWorkflow(t *testing.T) {
	workflowReader := loadFixture(t, "test_workflow.yml")
	assertion := assert.New(t)

	parser := Parser{}
	var (
		workflow types.Workflow
		err      error
	)

	workflow, err = parser.ParseWorkflow(workflowReader)

	assertion.Nil(err)
	assertion.NotNil(workflow.BaseDefinition)
	assertion.Equal("testid", workflow.ID)
	assertion.Equal("testkind", workflow.Kind)
	assertion.Equal("testname", workflow.MetaData.Name)
	assertion.Equal("testversion", workflow.MetaData.Version)
	assertion.Equal("testdescription", workflow.MetaData.Description)

	err = workflowReader.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestParserParseWorkflowInvalidReader(t *testing.T) {
	invalidReader := InvalidReader{}

	parser := Parser{}
	_, err := parser.ParseWorkflow(invalidReader)

	assert.Equal(t, errInvalidReader, err)
}

func TestParserParseWorkflowInvalidYaml(t *testing.T) {
	workflowReader := loadFixture(t, "invalid.txt")
	parser := Parser{}

	_, err := parser.ParseWorkflow(workflowReader)

	assert.NotNil(t, err)

	err = workflowReader.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestParserParseAction(t *testing.T) {
	workflowReader := loadFixture(t, "test_action.yml")
	assertion := assert.New(t)

	parser := Parser{}
	var (
		workflow types.Action
		err      error
	)

	workflow, err = parser.ParseAction(workflowReader)

	assertion.Nil(err)
	assertion.NotNil(workflow.BaseDefinition)
	assertion.Equal("testid", workflow.ID)
	assertion.Equal("testkind", workflow.Kind)
	assertion.Equal("testname", workflow.MetaData.Name)
	assertion.Equal("testversion", workflow.MetaData.Version)
	assertion.Equal("testdescription", workflow.MetaData.Description)

	err = workflowReader.Close()
	if err != nil {
		t.Fatal(err)
	}
}

func TestParserParseActionInvalidReader(t *testing.T) {
	invalidReader := InvalidReader{}

	parser := Parser{}
	_, err := parser.ParseAction(invalidReader)

	assert.Equal(t, errInvalidReader, err)
}

func TestParserParseActionInvalidYaml(t *testing.T) {
	workflowReader := loadFixture(t, "invalid.txt")
	parser := Parser{}

	_, err := parser.ParseAction(workflowReader)

	assert.NotNil(t, err)

	err = workflowReader.Close()
	if err != nil {
		t.Fatal(err)
	}
}
