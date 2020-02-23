package types

import (
	"errors"
)

// Kinds is the type for the allowed kind list
type Kinds []string

// IsAllowed checks if a kind is allowed
func (kinds Kinds) IsAllowed(kind string) bool {
	for _, kindItem := range kinds {
		if kindItem == kind {
			return true
		}
	}

	return false
}

// AllowedKinds defines all valid kinds
var AllowedKinds = Kinds{
	WorkflowType,
	ActionType,
}

var (
	// ErrInvalidID is returned for invalid ids
	ErrInvalidID = errors.New("invalid id")
	// ErrInvalidKind is returned for invalid a invalid kind
	ErrInvalidKind = errors.New("invalid kind")
)

// BaseDefinition is the base of all definitions
type BaseDefinition struct {
	ID       string   `yaml:"id"`
	Kind     string   `yaml:"kind"`
	MetaData MetaData `yaml:"metadata"`
}

// MetaData hold the metadata for all definitions
type MetaData struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
}

// Validate validates the base definition
func (baseDefinition BaseDefinition) Validate() error {
	if baseDefinition.ID == "" {
		return ErrInvalidID
	}

	if !AllowedKinds.IsAllowed(baseDefinition.Kind) {
		return ErrInvalidKind
	}

	return nil
}
