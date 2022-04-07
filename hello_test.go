package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello("Go for it man")
	expected := "Go for it man"

	if got != expected {
		t.Errorf("Failed Got: %s, Expected %s.", got, expected)
	} else {
		t.Logf("Succces Got: %s, Expected %s.", got, expected)
	}
}