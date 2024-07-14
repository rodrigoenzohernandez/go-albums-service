package main_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockServer struct {
	mock.Mock
}

func (m *MockServer) InitServer() {
	m.Called()
}

type Server interface {
	InitServer()
}

func StartApplication(server Server) {
	server.InitServer()
}

func TestMainFunction(t *testing.T) {
	testServer := new(MockServer)
	testServer.On("InitServer").Return()

	StartApplication(testServer)

	testServer.AssertExpectations(t)
}
