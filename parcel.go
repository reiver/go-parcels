package parcels


import (
	"io"
)


type Parcel interface {
	Bytes() []byte
	Reader() io.Reader
	Runes() []rune
	String() string
	WriteTo(w io.Writer) (int64, error)

	Then(func(Parcel)Parcel) Parcel
}

