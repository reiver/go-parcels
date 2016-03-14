package parcels


import (
	"github.com/reiver/go-oi"

	"io"
	"unicode/utf8"
)


func (parcel internalRuneParcel) WriteTo(w io.Writer) (int64, error) {

	var buffer [utf8.UTFMax]byte
	p := buffer[:]

	size := utf8.EncodeRune(p, parcel.content)

	p = p[:size]

	return oi.LongWrite(w, p)
}
