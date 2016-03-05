package parcels


import (
	"io"
	"strings"
)


type internalStringParcel struct {
	content string
}


// FromString returns a new Parcel whose contents is the value of the string.
//
// Example
//
//	parcel := parcels.FromString("This is some text.")
//	
//	asBytes  := parcel.Bytes()
//	asReader := parcel.Reader()
//	asRunes  := parcel.Runes()
//	asString := parcel.String()
func FromString(content string) Parcel {

	parcel := internalStringParcel{
		content:content,
	}

	return &parcel
}


func (parcel *internalStringParcel) Bytes() []byte {
	return []byte(parcel.content)
}


func (parcel *internalStringParcel) Reader() io.Reader {
	return strings.NewReader(parcel.content)
}


func (parcel *internalStringParcel) Runes() []rune {
	return []rune(parcel.content)
}


func (parcel *internalStringParcel) String() string {
	return parcel.content
}


func (parcel *internalStringParcel) Then(fn func(Parcel)Parcel) Parcel {
        return fn(parcel)
}

