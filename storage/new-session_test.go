package storage

import (
	"testing"
)

func TestNewSession(t *testing.T) {
	s, err := NewSession()
	if err != nil {
		t.Error("Error with creating session: ", err)
	}

	defer s.Close()
}
