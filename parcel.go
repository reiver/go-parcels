package parcels


import (
	"io"
)


type Parcel interface {
	Bytes() []byte
	Reader() io.Reader
	Runes() []rune
	String() string

	Then(func(Parcel)Parcel) Parcel
}

