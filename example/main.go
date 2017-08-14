package main

import (
	"readline"
)

func main() {
	prompt := "test> "
	scanner := &readline.Scanner{}
	scanner.Register(new(Test1))
	for scanner.Scan(prompt) {
		line := scanner.Text()
		if len(line) > 0 {
			scanner.Callback("Test1", line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
