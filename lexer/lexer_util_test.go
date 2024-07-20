package lexer

import (
	"testing"
)

func TestSkipWhiteSpaceUtil(t *testing.T) {
	input := ` ;a`
	l := New(input)

	l.skipWhiteSpace()
	if l.ch != ';' {
		t.Fatalf("test failed expected=%q, got=%q", 'a', l.ch)
	}
}
