package saver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	f := NewFact()
	s := NewMockedSaver()
	b := "bucket"
	assert.NoError(t, s.Save(f, b))
	assert.Contains(t, s.facts[b], f)
}

func TestNewMockedSaver(t *testing.T) {
	s := NewMockedSaver()
	assert.Equal(t, make(map[string][]*Fact), s.facts)
}
