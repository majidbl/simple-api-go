package util

import (
	"testing"
)

func TestUtil(t *testing.T) {
	pid, perr := GetPersonalId("teacher")
	if perr != nil {
		t.Fatalf("Expected no error, got %v", perr)
	}
	c := pid[len(pid)-2:]
	if c == "01" {
		t.Errorf("got end with %s,but pid must be ended with %s", pid, c)
	}
}
