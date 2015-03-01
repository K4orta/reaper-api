package storage

import (
	"testing"
)

func TestNewSession(t *testing.T) {
	s, err := CreateConnection()
	if err != nil {
		t.Error("Error with creating session: ", err)
	}

	defer s.Close()
}
