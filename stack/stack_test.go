package stack_test

import (
	"testing"

	"github.com/ignaciocon/dsa/stack"
)

func TestStack(t *testing.T) {
	newStack := stack.Stack{}

	if &newStack == nil {
		t.Error(`Error`)
	}
}
