package models

import "errors"

type Errors struct {
}

var (
	ErrNoRecord = errors.New("models: no matching record found")
)

func (e *Errors) NoRecordError() error {
	return ErrNoRecord
}
