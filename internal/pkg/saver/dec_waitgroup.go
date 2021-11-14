package saver

import (
	"sync"
)

type WithWaitGroupSaver struct {
	Saver
	WaitGroup *sync.WaitGroup
}

func NewWithWaitGroupSaver(saver Saver, wg *sync.WaitGroup) *WithWaitGroupSaver {
	return &WithWaitGroupSaver{
		Saver:     saver,
		WaitGroup: wg,
	}
}

func (s *WithWaitGroupSaver) Save(f *Fact, b string) error {
	defer s.WaitGroup.Done()
	return s.Saver.Save(f, b)
}
