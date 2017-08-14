package main

import (
	"readline"
)

func main() {
	prompt := "test> "
	scanner := &readline.Scanner{}
	for scanner.Scan(prompt) {
		line := scanner.Text()
		if len(line) > 0 {
			scanner.Callback("test", line)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
