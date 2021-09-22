package server

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockedServer(t *testing.T) {
	tests := []struct {
		name       string
		mockReturn error
	}{
		{
			"returning nil",
			nil,
		},
		{
			"returning error",
			errors.New("test error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMockedServer(func() error {
				return tt.mockReturn
			})
			if assert.IsType(t, &MockedServer{}, m) {
				r := m.startMock()
				assert.Equal(t, tt.mockReturn, r)
			}
		})
	}
}

func TestMockedServer_Start(t *testing.T) {
	tests := []struct {
		name       string
		mockReturn error
	}{
		{
			"returning nil",
			nil,
		},
		{
			"returning error",
			errors.New("test error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockedServer{
				startMock: func() error {
					return tt.mockReturn
				},
			}
			r := m.Start()
			assert.Equal(t, tt.mockReturn, r)
		})
	}
}
