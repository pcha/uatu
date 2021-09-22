package server

type MockedServer struct {
	startMock func() error
}

func NewMockedServer(startMock func() error) *MockedServer {
	return &MockedServer{startMock: startMock}
}

func (s *MockedServer) Start() error {
	return s.startMock()
}
