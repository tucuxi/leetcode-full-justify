package main

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	lines := fullJustify(words, 16)
	expected := []string{"This    is    an", "example  of text", "justification.  "}
	if !reflect.DeepEqual(expected, lines) {
		t.Fatalf("expected %v, got %v", expected, lines)
	}
}
