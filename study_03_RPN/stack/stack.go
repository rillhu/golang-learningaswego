package stack

import (
	"fmt"
	"strconv"
)

const (
	LEN = 12
)

type MyStack struct {
	i    int
	data [LEN]int
}

func (s *MyStack) isFull() bool {
	if s.i > LEN-1 {
		return true
	} else {
		return false
	}
}

func (s *MyStack) isEmpty() bool {
	if s.i == 0 {
		return true
	} else {
		return false
	}
}

func (s *MyStack) Top() int {
	if s.isEmpty() == true{
		fmt.Println("No top value")

	}
	id := s.i - 1
	return s.data[id]
}

func (s *MyStack) Push(k int) {
	if s.isFull() == true {
		fmt.Println("Stack is full")
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *MyStack) Pop() (k int) {
	if s.isEmpty() == true {
		fmt.Println("Stack is empty")
		return
	}
	s.i--
	tmp := s.data[s.i]
	//s.data[s.i] = 0
	return tmp
}

func (s *MyStack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

