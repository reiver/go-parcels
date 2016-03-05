# go-parcels

A library that provides a convenient immutable way of passing around a piece of data, with easy ways of getting that data as
`[]byte`, `io.Reader`, `[]rune`, or `string` types, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-parcels

[![GoDoc](https://godoc.org/github.com/reiver/go-parcels?status.svg)](https://godoc.org/github.com/reiver/go-parcels)


## Example
```
import (
	"github.com/reiver/go-parcels"
)

// ...

parcel := parcels.FromString("This is some content!")

asBytes  := parcel.Bytes()
asReader := parcel.Reader()
asRunes  := parcel.Runes()
asString := parcel.String()
```
