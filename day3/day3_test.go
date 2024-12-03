package day3

import (
	"reflect"
	"testing"
)

func TestFindCommands(t *testing.T) {
	testInput := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	got := FindCommands(testInput)
	want := []NumPair{
		{2, 4},
		{5, 5},
		{11, 8},
		{8, 5},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestReduceCommands(t *testing.T) {
	testInput := []NumPair{
		{2, 4},
		{5, 5},
		{11, 8},
		{8, 5},
	}

	got := ReduceCommands(testInput)
	want := 161

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
