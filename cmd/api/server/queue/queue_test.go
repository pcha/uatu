package queue

import (
	"sync"
	"testing"

	"the-one/internal/saver"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := saver.NewMockedSaver()
	q := New(s)
	assert.NotNil(t, q)
	assert.Same(t, s, q.saver)
}

func TestNewLoggeable(t *testing.T) {
	l := saver.NewFact()
	assert.NotNil(t, l)
}

func TestQueue_Receive(t *testing.T) {
	q := New(nil)
	b := "bucket"
	d := saver.NewFact()
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
	q := New(saver.NewWithWaitGroupSaver(saver.NewMockedSaver(), &wg))
	wg.Add(3)
	q.StartListening()
	q.Receive("bucket", saver.NewFact())
	q.Receive("bucket", saver.NewFact())
	q.Receive("bucket", saver.NewFact())
	wg.Wait()
	q.StopListening()
}
