# readline
Pure Go readline package

## Example

main.go like this:
```go
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
```

and test.go like this:
```go
type Test1 int

func (this *Test1) Echo(args readline.Args, reply *readline.Reply) error {
    if len(args.A) > 0 {
        fmt.Println(args.A[0])
    }
    return nil
}
```
