package parcels


import (
	"bytes"
	"io"
)


type internalBytesParcel struct {
	content []byte
}


// FromBytes returns a new Parcel whose contents is the value of the []byte.
// Note that the contents the input []byte 'b' is copied, and the new Parcel
// is not attached to the original backing memory.
//
// Example
//
//      parcel := parcels.FromBytes( []byte{72, 69, 76, 76, 79} )
//      
//      asBytes  := parcel.Bytes()
//      asReader := parcel.Reader()
//      asRunes  := parcel.Runes()
//      asString := parcel.String()
func FromBytes(b []byte) Parcel {

	content := append([]byte(nil), b...)

	parcel := internalBytesParcel{
		content:content,
	}

	return &parcel
}


func (parcel *internalBytesParcel) Bytes() []byte {
	return append([]byte(nil), parcel.content...)
}


func (parcel *internalBytesParcel) Reader() io.Reader {
	return bytes.NewReader(parcel.content)
}


func (parcel *internalBytesParcel) Runes() []rune {
	return []rune(string(parcel.content))
}


func (parcel *internalBytesParcel) String() string {
	return string(parcel.content)
}


func (parcel *internalBytesParcel) Map(fn func(Parcel)Parcel) Parcel {
	return fn(parcel)
}
