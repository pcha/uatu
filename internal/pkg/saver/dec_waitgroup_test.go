package saver

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWithWaitGroupSaver(t *testing.T) {
	var wg sync.WaitGroup
	s := NewWithWaitGroupSaver(NewMockedSaver(), &wg)
	assert.NotNil(t, s)
}

func TestWithWaitGroupSaver_Save(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	s := NewWithWaitGroupSaver(NewMockedSaver(), &wg)
	f := NewFact()
	assert.NoError(t, s.Save(f, "bucket"))
	wg.Wait()
}
