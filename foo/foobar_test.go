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

func TestGivenTwoWantTwo(t *testing.T)()  {
	given := 2
	want := "2"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}

func TestGivenThreeWantThree(t *testing.T)()  {
	given := 3
	want := "Foo"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}

func TestGivenFourWantFour(t *testing.T)()  {
	given := 4
	want := "4"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}

func TestGivenFiveWantBar(t *testing.T)()  {
	given := 5
	want := "Bar"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}

func TestGivenSixWantSix(t *testing.T)()  {
	given := 6
	want := "Foo"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}

func TestGivenSevenWantSeven(t *testing.T)()  {
	given := 7
	want := "7"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}

func TestGivenEightWantEight(t *testing.T)()  {
	given := 8
	want := "8"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}

func TestGivenNineWantFoo(t *testing.T)()  {
	given := 9
	want := "Foo"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}

func TestGivenTenWantBar(t *testing.T)()  {
	given := 10
	want := "Bar"

	result := foo.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q, want %s", given, result, want)
	}
}