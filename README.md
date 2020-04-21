# number-to-words

## Next plans
- add more test cases
- clean code
  - remove duplicate code
  - handle error
- improve speed
- new features
  - convert without baht text
  - new language support(?)

## Install
```go get -u github.com/sujamess/number-to-word```

## Examples

``` Go
func main() {
  num := 555.55

  text, err := numbertowords.Convert(num, "en")
  if err != nil {
        fmt.Println("Error occurred: %v", err)
  }
  
  fmt.Println(text) // result: fiftyhundred fifty five point five five baht
}
```

## Benchmarks
```
pkg: number-to-words
BenchmarkConvertEN-8   	  610806	      1660 ns/op	     536 B/op	      22 allocs/op
BenchmarkConvertTH-8   	  908407	      1318 ns/op	     512 B/op	      15 allocs/op
```