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
        fmt.Println("Error occured: %v", err)
  }
  
  fmt.Println(text) // result: fiftyhundred fifty five point five five baht
}
```

## Benchmarks
```
pkg: number-to-words
BenchmarkIntegerEN-8   	  431254	      2525 ns/op	     896 B/op	      36 allocs/op
BenchmarkDecimalEN-8   	 5162589	       234 ns/op	     128 B/op	       4 allocs/op
BenchmarkConvertEN-8   	  744618	      1657 ns/op	     536 B/op	      22 allocs/op
BenchmarkReaderTH-8    	  483110	      2390 ns/op	     946 B/op	      30 allocs/op
BenchmarkDecimalTH-8   	 2538355	       486 ns/op	     192 B/op	       7 allocs/op
BenchmarkConvertTH-8   	  890694	      1311 ns/op	     512 B/op	      15 allocs/op
```