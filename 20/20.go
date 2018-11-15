package leetcode_20

import "fmt"

type Stack struct {
	Arr []string
}

func (s *Stack) pop() string {
	var result string
	if len(s.Arr) == 0 {
		return ""
	}
	result = s.Arr[len(s.Arr)-1]
	s.Arr = s.Arr[:len(s.Arr)-1]
	return result
}

func (s *Stack) push(c string) {
	s.Arr = append(s.Arr, c)
}

func (s *Stack) print() {
	for _, v := range s.Arr {
		fmt.Println(v)
	}
}
func isValid(s string) bool {
	stack := Stack{}
	for _, v := range s {
		cur := string(v)
		if cur == "(" || cur == "[" || cur == "{" {
			stack.push(cur)
		} else {
			tmp := stack.pop()
			if (cur == ")" && tmp == "(") || (cur == "]" && tmp == "[") || (cur == "}" && tmp == "{") {
				continue
			} else {
				return false
			}
		}
	}

	if len(stack.Arr) != 0 {
		return false
	}
	return true
}
