# number-to-words

## Next plans
- clean code
  - remove duplicate code
  - handle error
- improve speed
- new features
  - convert without baht text
  - new language(?)

## Install
```go get -u github.com/sujamess/number-to-word```

## Examples

```
func main() {
        num := 555.55
  
        text, err := numbertowords.Convert(num, "en")
        if err != nil {
              fmt.Println("Error occured: %v", err)
        }
        
        fmt.Println(text) // result: fiftyhundred fifty five point five five baht
}
```
