package parcels


import (
	"github.com/reiver/go-oi/test"

	"bytes"
	"io/ioutil"
	"strings"

	"testing"
)


func TestParcelFromRune(t *testing.T) {

	tests := []struct{
		Rune rune
	}{
		{
			Rune: 'A',
		},
		{
			Rune: 'B',
		},
		{
			Rune: 'C',
		},
		{
			Rune: 'D',
		},
		{
			Rune: 'E',
		},



		{
			Rune: 'a',
		},
		{
			Rune: 'b',
		},
		{
			Rune: 'c',
		},
		{
			Rune: 'd',
		},
		{
			Rune: 'e',
		},



		{
			Rune: '0',
		},
		{
			Rune: '1',
		},
		{
			Rune: '2',
		},
		{
			Rune: '3',
		},
		{
			Rune: '4',
		},
		{
			Rune: '5',
		},
		{
			Rune: '6',
		},
		{
			Rune: '7',
		},
		{
			Rune: '8',
		},
		{
			Rune: '9',
		},



		{
			Rune: '٠', // Arabic-Indic Digits
		},
		{
			Rune: '١', // Arabic-Indic Digits
		},
		{
			Rune: '٢', // Arabic-Indic Digits
		},
		{
			Rune: '٣', // Arabic-Indic Digits
		},
		{
			Rune: '٤', // Arabic-Indic Digits
		},
		{
			Rune: '٥', // Arabic-Indic Digits
		},
		{
			Rune: '٦', // Arabic-Indic Digits
		},
		{
			Rune: '٧', // Arabic-Indic Digits
		},
		{
			Rune: '٨', // Arabic-Indic Digits
		},
		{
			Rune: '٩', // Arabic-Indic Digits
		},



		{
			Rune: '۰', // Extended Arabic-Indic Digits
		},
		{
			Rune: '۱', // Extended Arabic-Indic Digits
		},
		{
			Rune: '۲', // Extended Arabic-Indic Digits
		},
		{
			Rune: '۳', // Extended Arabic-Indic Digits
		},
		{
			Rune: '۴', // Extended Arabic-Indic Digits
		},
		{
			Rune: '۵', // Extended Arabic-Indic Digits
		},
		{
			Rune: '۶', // Extended Arabic-Indic Digits
		},
		{
			Rune: '۷', // Extended Arabic-Indic Digits
		},
		{
			Rune: '۸', // Extended Arabic-Indic Digits
		},
		{
			Rune: '۹', // Extended Arabic-Indic Digits
		},



		{
			Rune: 'Ⅰ',
		},
		{
			Rune: 'Ⅱ',
		},
		{
			Rune: 'Ⅲ',
		},
		{
			Rune: 'Ⅳ',
		},
		{
			Rune: 'Ⅴ',
		},
		{
			Rune: 'Ⅵ',
		},
		{
			Rune: 'Ⅶ',
		},
		{
			Rune: 'Ⅷ',
		},
		{
			Rune: 'Ⅸ',
		},
		{
			Rune: 'Ⅹ',
		},
		{
			Rune: 'Ⅺ',
		},
		{
			Rune: 'Ⅻ',
		},
		{
			Rune: 'Ⅼ',
		},
		{
			Rune: 'Ⅽ',
		},
		{
			Rune: 'Ⅾ',
		},
		{
			Rune: 'Ⅿ',
		},



		{
			Rune: 'ⅰ',
		},
		{
			Rune: 'ⅱ',
		},
		{
			Rune: 'ⅲ',
		},
		{
			Rune: 'ⅳ',
		},
		{
			Rune: 'ⅴ',
		},
		{
			Rune: 'ⅵ',
		},
		{
			Rune: 'ⅶ',
		},
		{
			Rune: 'ⅷ',
		},
		{
			Rune: 'ⅸ',
		},
		{
			Rune: 'ⅹ',
		},
		{
			Rune: 'ⅺ',
		},
		{
			Rune: 'ⅻ',
		},
		{
			Rune: 'ⅼ',
		},
		{
			Rune: 'ⅽ',
		},
		{
			Rune: 'ⅾ',
		},
		{
			Rune: 'ⅿ',
		},



		{
			Rune: 'ↀ',
		},
		{
			Rune: 'ↁ',
		},
		{
			Rune: 'ↂ',
		},
		{
			Rune: 'Ↄ',
		},
		{
			Rune: 'ↄ',
		},
		{
			Rune: 'ↅ',
		},
		{
			Rune: 'ↆ',
		},
		{
			Rune: 'ↇ',
		},
		{
			Rune: 'ↈ',
		},
	}


	for testNumber, test := range tests {

		parcel := ParcelFromRune(test.Rune)


		if expected, actual := string(test.Rune), parcel.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}


		if expected, actual := string(test.Rune), string(parcel.Bytes()); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}


		if expected, actual := string(test.Rune), string(parcel.Runes()); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}


		actualReaderBytes, err := ioutil.ReadAll(parcel.Reader())
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v; for %q.", testNumber, err, string(test.Rune))
			continue
		}
		if expected, actual := string(test.Rune), string(actualReaderBytes); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}


		var buffer bytes.Buffer
		n, err := parcel.WriteTo(&buffer)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v; for %q.", testNumber, err, string(test.Rune))
			continue
		}
		if expected, actual := int64(len([]byte(string(test.Rune)))), n; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for %q.", testNumber, expected, actual, string(test.Rune))
			continue
		}
		if expected, actual := string(test.Rune), buffer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q", testNumber, expected, actual)
			continue
		}


		var writer oitest.ShortWriter
		n, err = parcel.WriteTo(&writer)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v; for %q.", testNumber, err, string(test.Rune))
			continue
		}
		if expected, actual := int64(len([]byte(string(test.Rune)))), n; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for %q.", testNumber, expected, actual, string(test.Rune))
			continue
		}
		if expected, actual := string(test.Rune), writer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q", testNumber, expected, actual)
			continue
		}
	}
}


func TestRuneParcelThen(t *testing.T) {


	tests := []struct{
		Rune    rune
		Thens []struct {
			Func func(Parcel)Parcel
			Expected string
		}
	}{
		{
			Rune: 'A',
			Thens: []struct{
				Func func(Parcel)Parcel
				Expected string
			}{
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToUpper(parcel.String())
						return ParcelFromString(s)
					},
					Expected: "A",
				},
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToLower(parcel.String())
						return ParcelFromString(s)
					},
					Expected: "a",
				},
			},
		},



		{
			Rune: 'a',
			Thens: []struct{
				Func func(Parcel)Parcel
				Expected string
			}{
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToUpper(parcel.String())
						return ParcelFromString(s)
					},
					Expected: "A",
				},
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToLower(parcel.String())
						return ParcelFromString(s)
					},
					Expected: "a",
				},
			},
		},
	}


	for testNumber, test := range tests {

		parcel := ParcelFromRune(test.Rune)

		for thenNumber, then := range test.Thens {

			if expected, actual := then.Expected, parcel.Then(then.Func).String(); expected != actual {
				t.Errorf("For test #%d and then #%d, expected %q, but actually got %q.", testNumber, thenNumber, expected, actual)
				continue
			}
		}
	}
}
