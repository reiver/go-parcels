package parcels


import (
	"bytes"
	"io"
	"unicode/utf8"
)


type internalRuneParcel struct {
	content rune
}


// ParcelFromRune returns a new Parcel whose contents is the value of the rune.
//
// Example
//
//	parcel := parcels.ParcelFromRune('۵')
//	
//	asBytes  := parcel.Bytes()
//	asReader := parcel.Reader()
//	asRunes  := parcel.Runes()
//	asString := parcel.String()
//	
//	n, err := parcel.WriteTo(writer)
func ParcelFromRune(r rune) Parcel {

	parcel := internalRuneParcel{
		content:r,
	}

	return parcel
}


func (parcel internalRuneParcel) Bytes() []byte {
	var buffer [utf8.UTFMax]byte
	p := buffer[:]

	size := utf8.EncodeRune(p, parcel.content)

	return append([]byte(nil), p[:size]...)
}


func (parcel internalRuneParcel) Reader() io.Reader {
	var buffer [utf8.UTFMax]byte
	p := buffer[:]

	size := utf8.EncodeRune(p, parcel.content)

	p = p[:size]

	return bytes.NewReader(p)
}


func (parcel internalRuneParcel) Runes() []rune {
	return []rune{parcel.content}
}


func (parcel internalRuneParcel) String() string {
	return string(parcel.content)
}


func (parcel internalRuneParcel) Then(fn func(Parcel)Parcel) Parcel {
	return fn(parcel)
}
