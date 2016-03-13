package parcels


import (
	"github.com/reiver/go-oi"

	"io"
)


func (parcel *internalBytesParcel) WriteTo(w io.Writer) (int64, error) {

	if len(parcel.content) < lenBuffer {
		return parcel.smallWriteTo(w)
	} else {
		return parcel.bigWriteTo(w)
	}
}


func (parcel *internalBytesParcel) smallWriteTo(w io.Writer) (int64, error) {

	var buffer [lenBuffer]byte
	p := buffer[:]


	// Copy the data from the parcel's content over to a buffer so that the
	// writer does not have an opportunity to write to parcel.conetent (so
	// that we can conceptually keep the parcel immutable).
	copy(p, parcel.content)
	p = p[:len(parcel.content)]


	return oi.LongWrite(w, p)
}


func (parcel *internalBytesParcel) bigWriteTo(w io.Writer) (int64, error) {

	var buffer [lenBuffer]byte
	p := buffer[:]


	lenContent := len(parcel.content)


	numWritten := int64(0)
	for iteration:=0; true; iteration++ {
		index := iteration * lenBuffer

		if !(index < lenContent) {
			break
		}

		// Copy the data from the parcel's content over to a buffer so that the
		// writer does not have an opportunity to write to parcel.conetent (so
		// that we can conceptually keep the parcel immutable).
		pcc := parcel.content[index:]
		copy(p, pcc)
		if lenPCC := len(pcc); lenPCC < lenBuffer {
			p = p[:lenPCC]
		}

		if len(p) < 1 {
			break
		}

		n, err := oi.LongWrite(w, p)
		numWritten += n
		if nil != err {
			return numWritten, err
		}
	}


	return numWritten, nil
}
