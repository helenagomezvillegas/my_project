package main

import (
	"reflect"
	"testing"
)

func TestMove(t *testing.T) {
	commands := []string{"R"}

	got := Moves(commands)
	want := []int{1, 0}

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("got %d want %d given, %v", got, want, commands)

	}
}
