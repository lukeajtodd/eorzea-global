# Pointers

Go is fully garbage collected. It has full support for pointers and pointer arithmetic.

```go

func getAnswer () (answer *int) {
  x := 42
  return &x
}

func main () {
  answers := *getAnswer()
  fmt.Println("Answer is:", answer)
}
```
