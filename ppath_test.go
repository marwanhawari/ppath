package main

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	tests := []struct {
		array    []string
		target   string
		expected bool
	}{
		{[]string{"a", "b"}, "a", true},
		{[]string{"a", "b"}, "b", true},
		{[]string{"a", "b"}, "c", false},
	}
	for _, test := range tests {
		got := contains(test.array, test.target)
		if got != test.expected {
			t.Errorf("contains(%v, %v) == %v, want %v", test.array, test.target, got, test.expected)
		}
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		array    []string
		expected []string
	}{
		{[]string{"a", "b", "a"}, []string{"a", "b"}},
		{[]string{"a", "b", "b"}, []string{"a", "b"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
	}
	for _, test := range tests {
		got := unique(test.array)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("unique(%v) == %v, want %v", test.array, got, test.expected)
		}
	}
}
