package datastructure

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
