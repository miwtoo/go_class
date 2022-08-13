package foo_test

import (
	"testing"

	"github.com/miwtoo/go_class/foo"
)

// cd foo
// go test
func TestGivenOneWantOne(t *testing.T)()  {
	given := 1
	want := "1"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}