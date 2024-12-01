package day1

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	got1, got2, err := ParseInput(input)
	want1 := []int{3, 4, 2, 1, 3, 3}
	want2 := []int{4, 3, 5, 3, 9, 3}

	if err != nil {
		t.Errorf("input created an error: %v", err)
	}

	if !reflect.DeepEqual(got1, want1) {
		t.Errorf("got %v, want %v", got1, want1)
	}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %v, want %v", got2, want2)
	}
}

func TestGetTotal(t *testing.T) {
	input := `3   4
	4   3
	2   5
	1   3
	3   9
	3   3`

	got := GetTotal(input)
	want := 11

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
