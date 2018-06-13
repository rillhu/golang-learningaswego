package main

import (
	"fmt"
)

type I interface {
	Get() int
	Put(int)
}

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}

func (p *S) Put(v int) {
	p.i = v
}

/*type R also implements the interface I*/
type R struct {
	i int
}

func (p *R) Get() int {
	return p.i
}

func (p *R) Put(v int) {
	p.i = v
}

func f(p I) {

	switch t := p.(type) {
	case *S:
		fmt.Println(t)
		p.Put(1)
		fmt.Println(p.Get())
	case *R:
		fmt.Println(t)
		p.Put(2)
		fmt.Println(p.Get())
	default:

	}
	p.Put(1)

}

func g(something interface{}) int {
	return something.(I).Get()
}

func main() {

	//A is a struct pointer
	//A implement the interface I
	A := new(S)
	A.Put(9)
	fmt.Println(A.Get())
	//So A can be assigned to interface I variable
	//Use pointer is due to the implementation is implemented with pointer func (s *S)
	f(A)

	/*type R also implements the interface I*/
	B := new(R)
	f(B)

	var B2 R
	f(&B2)

	A2 := new(S)
	fmt.Println(g(A2)) //Print 0
}
