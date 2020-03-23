package main

import "fmt"

type people struct{}

func (p *people) showA() {
	fmt.Println("i am people showA")
	p.showB()
}
func (p *people) showB() {
	fmt.Println("i am people showB")
}

type student struct {
	people
}

func (t *student) showB() {
	fmt.Println("i am student showB")
}
func main() {
	t := student{}
	t.showA()
	t.showB()
}

// i am people showA
// i am people showB 接受者是p，所以调用的是p的方法
