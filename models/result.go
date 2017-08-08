package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/pop/slices"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Result struct {
	ID           uuid.UUID     `json:"id" db:"id"`
	CreatedAt    time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" db:"updated_at"`
	Num          int           `json:"num" db:"num"`
	Line         string        `json:"line" db:"line"`
	Matches      slices.String `json:"matches" db:"matches"`
	ExpressionID uuid.UUID     `json:"expression_id" db:"expression_id"`
}

// String is not required by pop and may be deleted
func (r Result) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Results is not required by pop and may be deleted
type Results []Result

// String is not required by pop and may be deleted
func (r Results) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (r *Result) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: r.Num, Name: "Num"},
		&validators.StringIsPresent{Field: r.Line, Name: "Line"},
	), nil
}

// ValidateSave gets run every time you call "pop.ValidateSave" method.
// This method is not required and may be deleted.
func (r *Result) ValidateSave(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateUpdate" method.
// This method is not required and may be deleted.
func (r *Result) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
