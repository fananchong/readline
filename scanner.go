package readline

import (
	"fmt"
)

type Scanner struct {
	line string
	err  error
}

func (this *Scanner) Scan(prompt string) bool {

	fmt.Print(prompt)

	var c byte
	var err error
	var b []byte
	for {
		_, err = fmt.Scanf("%c", &c)
		if err == nil && c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}
	this.line = string(b)
	this.err = err
	return err == nil
}

func (this *Scanner) Text() string {
	return this.line
}

func (this *Scanner) Err() error {
	return this.err
}

func (this *Scanner) Callback(clas string, cmd string) {

}
