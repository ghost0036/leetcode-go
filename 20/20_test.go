package leetcode_20

import (
	"fmt"
	"testing"
)

func TestNormal(t *testing.T) {
	fmt.Println(isValid("([)]"))
}

func TestStack(t *testing.T) {
	stack := Stack{}
	stack.print()

	stack.push("a")
	stack.push("b")
	stack.push("c")

	stack.print()

	stack.pop()

	stack.print()

}
