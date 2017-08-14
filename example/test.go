package main

import (
	"fmt"
	"readline"
)

type Test1 int

func (this *Test1) Echo(args readline.Args, reply *readline.Reply) error {
	if len(args.A) > 0 {
		fmt.Println(args.A[0])
	}
	return nil
}
