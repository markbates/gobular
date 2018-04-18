package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Expression struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Expression string    `json:"expression" db:"expression"`
	TestString string    `json:"test_string" db:"test_string"`
}

// String is not required by pop and may be deleted
func (e Expression) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Expressions is not required by pop and may be deleted
type Expressions []Expression

// String is not required by pop and may be deleted
func (e Expressions) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (e *Expression) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Expression, Name: "Expression"},
		&validators.StringIsPresent{Field: e.TestString, Name: "TestString"},
	), nil
}

// ValidateSave gets run every time you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (e *Expression) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (e *Expression) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
