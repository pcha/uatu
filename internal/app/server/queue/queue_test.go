package queue

import (
	"sync"
	"testing"

	saver2 "the-one/internal/pkg/saver"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := saver2.NewMockedSaver()
	q := New(s)
	assert.NotNil(t, q)
	assert.Same(t, s, q.saver)
}

func TestNewLoggeable(t *testing.T) {
	l := saver2.NewFact()
	assert.NotNil(t, l)
}

func TestQueue_Receive(t *testing.T) {
	q := New(nil)
	b := "bucket"
	d := saver2.NewFact()
	expected := Queueable{
		Bucket: b,
		Data:   d,
	}
	go q.Receive(b, d)
	got := <-q.channel
	assert.Equal(t, expected, got)
}

func TestQueue_StartListening(t *testing.T) {
	var wg sync.WaitGroup
	q := New(saver2.NewWithWaitGroupSaver(saver2.NewMockedSaver(), &wg))
	wg.Add(3)
	q.StartListening()
	q.Receive("bucket", saver2.NewFact())
	q.Receive("bucket", saver2.NewFact())
	q.Receive("bucket", saver2.NewFact())
	wg.Wait()
	q.StopListening()
}
