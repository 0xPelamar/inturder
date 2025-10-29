package helper

import (
	"errors"
	"fmt"
)

type MultiError struct {
	errors []error
}

func NewMultiError() *MultiError {
	return &MultiError{
		errors: []error{},
	}
}

func (e *MultiError) Add(err error) {
	e.errors = append(e.errors, err)
}

func (e *MultiError) ToError() error {
	if len(e.errors) == 0 {
		return nil
	}
	errString := fmt.Sprintf("%d errors occured:\n", len(e.errors))
	for _, err := range e.errors {
		errString = fmt.Sprintf("%s\t* %s\n", errString, err.Error())
	}
	return errors.New(errString)
}
