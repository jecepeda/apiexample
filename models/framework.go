package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Framework struct {
	ID          uuid.UUID `json:"id" db:"id" fake:"skip"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" fake:"skip"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" fake:"skip"`
	Title       string    `json:"title" db:"title" fake:"hipster_word"`
	Description string    `json:"description" db:"description" fake:"hipster_paragraph"`
}

// String is not required by pop and may be deleted
func (f Framework) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Frameworks is not required by pop and may be deleted
type Frameworks []Framework

// String is not required by pop and may be deleted
func (f Frameworks) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Framework) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: f.Title, Name: "Title"},
		&validators.StringIsPresent{Field: f.Description, Name: "Description"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Framework) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Framework) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
