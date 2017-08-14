package main

import (
	"fmt"
)

type Test int

func (this *Test) Echo(msg string) {
	fmt.Println(msg)
}

func (this *Test) f0() {
	fmt.Println("f0")
}

func (this *Test) f1(p1 string) {
	fmt.Print("f1 ")
	fmt.Println(p1)
}

func (this *Test) f2(p1 string, p2 int) {
	fmt.Print("f2 ")
	fmt.Println(p1, p2)
}
