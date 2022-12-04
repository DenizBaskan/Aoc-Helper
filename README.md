# Aoc-Helper
Advent of code helper used to make tedious tasks such as reading and parsing input easier, downloaded inputs are cached for use later.

## Usage

```go
var sess aoc.Session

err := sess.InitSessionFromFile("../session.txt")
if err != nil {
    panic(err)
}
	
// sess.InitSession("session here")

input, err := sess.RetrieveInput(2021, 8)
if err != nil {
    panic(err)
}
	
inputLines := input.Strings() // returns input converted to a string array

inputFloats, err := input.Floats() // returns input converted to a float array
if err != nil {
    panic(err)
}

inputInts, err := input.Ints() // returns input converted to an int array
if err != nil {
    panic(err)
}

inputRaw := input.Raw() // returns raw input (no modifications)
```
