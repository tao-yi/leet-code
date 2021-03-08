package main

import "testing"

func ValidParentheses(s string) bool {
	stack := []rune{}
	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) < 1 {
				return false
			}
			top := stack[len(stack)-1]
			if top != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{input: "()", valid: true},
		{input: "(]", valid: false},
		{input: "(([])", valid: false},
		{input: "", valid: true},
	}

	for _, test := range tests {
		output := ValidParentheses(test.input)
		if output != test.valid {
			t.Errorf("want %t, got %t", test.valid, output)
		}
	}
}
