# Functions

## Lambdas

Functions can be first class objects. They can be passed around in lambda form.

```go
aFunc := func() bool {
  true
}
```

## Named return values

They can also return named values so when returning with nothing, those values are pulled from scope.

This is a poor example but you can see it would return `hasApple` automagically.

```go
func apples(fruits []string) (hasApple bool) {
  _, has := Find(fruits, "apple") // Some find implementation
  hasApple = has
  return
}

func main() {
  items := []string{"apple", "banana"}
  hasApple := apples(fruits)
}
```

## Method implementations

Methods can be attached to a struct as can be seen in the [locks](./locks.md) file when attaching the `Add` method to the earlier struct.
