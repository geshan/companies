package db

import (
	"errors"
	"testing"
)

type MockDB struct {
	PingFunc  func() error
	CloseFunc func() error
}

func (m *MockDB) Ping() error  { return m.PingFunc() }
func (m *MockDB) Close() error { return m.CloseFunc() }

func TestService_Health_OK(t *testing.T) {
	mock := &MockDB{
		PingFunc: func() error { return nil },
	}
	s := &Service{DB: mock}
	if err := s.Health(); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestService_Health_Error(t *testing.T) {
	mock := &MockDB{
		PingFunc: func() error { return errors.New("db error") },
	}
	s := &Service{DB: mock}
	if err := s.Health(); err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestService_Close_OK(t *testing.T) {
	closed := false
	mock := &MockDB{
		CloseFunc: func() error { closed = true; return nil },
	}
	s := &Service{DB: mock}
	if err := s.Close(); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if !closed {
		t.Errorf("expected closed to be true")
	}
}

func TestService_Close_NilDB(t *testing.T) {
	s := &Service{DB: nil}
	if err := s.Close(); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
