package saver

type MockedSaver struct {
	facts  map[string][]*Fact
	params map[string]string
}

func NewMockedSaver() *MockedSaver {
	return &MockedSaver{
		facts: make(map[string][]*Fact),
	}
}

func NewMockedSaverWithParams(params map[string]string) *MockedSaver {
	return &MockedSaver{
		facts:  make(map[string][]*Fact),
		params: params,
	}
}

func (s *MockedSaver) Save(f *Fact, b string) error {
	s.facts[b] = append(s.facts[b], f)
	return nil
}
