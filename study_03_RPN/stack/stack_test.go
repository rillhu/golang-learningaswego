package stack

import "testing"

//var stackTest MyStack

func TestMyStack_Push(t *testing.T) {
	stackTest := new(MyStack)
	stackTest.Push(100)
	id := stackTest.i - 1
	if stackTest.data[id] != 100 {
		t.Log("Push stack fail!")
		t.Fail()
	}
}

func TestMyStack_Pop(t *testing.T) {
	stackTest := new(MyStack)
	stackTest.Push(100)
	val := stackTest.Pop()

	if val != 100 {
		t.Log("Pop stack fail")
		t.Fail()
	}
}
