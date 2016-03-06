# go-parcels

A library that provides a convenient immutable way of passing around a piece of data, with easy ways of getting that data as
`[]byte`, `io.Reader`, `[]rune`, or `string` types, for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-parcels

[![GoDoc](https://godoc.org/github.com/reiver/go-parcels?status.svg)](https://godoc.org/github.com/reiver/go-parcels)


## Example #1a
```
import (
	"github.com/reiver/go-parcels"
)

// ...

parcel := parcels.ParcelFromString("This is some content!")

asBytes  := parcel.Bytes()
asReader := parcel.Reader()
asRunes  := parcel.Runes()
asString := parcel.String()
```


## Example 1b
```
import (
	"github.com/reiver/go-parcels"
)

// ...

parcel := parcels.ParcelFromBytes( []byte{72, 69, 76, 76, 79} )

asBytes  := parcel.Bytes()
asReader := parcel.Reader()
asRunes  := parcel.Runes()
asString := parcel.String()
```


## Example #2a
```
import (
	"github.com/reiver/go-parcels"
)

// ...

parcel := parcels.ParcelFromString("This is some content!")

newParcel := parcel.Then(func(parcel Parcel)Parcel{
	s := strings.ToUpper(parcel.String)
	return parcels.ParcelFromString(s)
})
```


## Example #2b
```
import (
	"github.com/reiver/go-parcels"
)

// ...

parcel := parcels.ParcelFromBytes( []byte{72, 69, 76, 76, 79} )

newParcel := parcel.Then(func(parcel Parcel)Parcel{
	s := strings.ToUpper(parcel.String)
	return parcels.ParcelFromString(s)
})
```
