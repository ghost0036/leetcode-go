package leetcode_71

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := Stack{}
	list := strings.Split(path, "/")
	for _, v := range list {
		switch v {
		case "":
			continue
		case ".":
			continue
		case "..":
			stack.pop()
		default:
			stack.push(v)
		}
	}
	return stack.printPath()

}

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

func (s *Stack) printPath() (result string) {
	result = ""
	if len(s.Arr) == 0 {
		return "/"
	}

	for _, v := range s.Arr {
		result = result + "/" + v
	}
	return
}
