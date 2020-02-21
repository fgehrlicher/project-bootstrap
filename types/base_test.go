package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseDefinitionValidate(t *testing.T) {
	baseDefinition := BaseDefinition{
		ID: "id",
		Kind: WorkflowType,
	}

	err := baseDefinition.Validate()
	assert.Nil(t, err)
}

func TestBaseDefinitionValidateNoID(t *testing.T) {
	baseDefinition := BaseDefinition{
		Kind: WorkflowType,
	}

	err := baseDefinition.Validate()
	assert.Equal(t, errInvalidID, err)
}

func TestBaseDefinitionValidateNoKind(t *testing.T) {
	baseDefinition := BaseDefinition{
		ID: "testid",
	}

	err := baseDefinition.Validate()
	assert.Equal(t, errInvalidKind, err)
}

func TestBaseDefinitionValidateInvalidKind(t *testing.T) {
	baseDefinition := BaseDefinition{
		ID:   "testid",
		Kind: "invalidkind",
	}

	err := baseDefinition.Validate()
	assert.Equal(t, errInvalidKind, err)
}
