package types

import (
	"errors"
)

// AllowedKinds defines all valid kinds
var AllowedKinds = []string{
	WorkflowType,
	ActionType,
}

var (
	errInvalidID = errors.New("invalid id")
	errInvalidKind = errors.New("invalid kind")
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
		return errInvalidID
	}

	for _, kind := range AllowedKinds {
		if baseDefinition.Kind == kind {
			return nil
		}
	}

	return errInvalidKind
}
