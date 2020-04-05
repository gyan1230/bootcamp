package main

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		input          string
		expectedOutput string
	}{
		{"", ""},
		{"a", "a"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"Hello, 世界", "界世 ,olleH"},
	}
	for _, v := range cases {
		if op := reverse(v.input); op != v.expectedOutput {
			t.Errorf("incorrect op for %s, expected :%s, got :%s\n", v.input, v.expectedOutput, op)
		}
	}
}
