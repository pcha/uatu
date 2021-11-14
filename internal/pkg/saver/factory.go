package saver

import (
	"fmt"
)

type UnknownTypeError struct {
	givenType string
}

func (e *UnknownTypeError) Error() string {
	return fmt.Sprintf("Unknown saver type %q", e.givenType)
}

func NewSaver(saverType string, params map[string]string) (Saver, error) {
	switch saverType {
	case "mock":
		return NewMockedSaverWithParams(params), nil
	case "mongo":
	case "mongodb":
		return NewMongoSaver(params)
	}
	return nil, &UnknownTypeError{givenType: saverType}
}
