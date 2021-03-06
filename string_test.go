package parcels


import (
	"github.com/reiver/go-oi/test"

	"bytes"
	"io/ioutil"
	"strings"

	"testing"
)


func TestParcelFromString(t *testing.T) {

	tests := []struct{
		String string
	}{
		{
			String: "",
		},



		{
			String: "apple",
		},
		{
			String: "banana",
		},
		{
			String: "cherry",
		},



		{
			String: "Hello world!",
		},



		{
			String: "😁😂😃😄😅😆😉😊😋😌😍😏😒😓😔😖😘😚😜😝😞😠😡😢😣😤😥😨😩😪😫😭😰😱😲😳😵😷",
		},



		{
			String: "0123456789",
		},
		{
			String: "٠١٢٣٤٥٦٧٨٩", // Arabic-Indic Digits
		},
		{
			String: "۰۱۲۳۴۵۶۷۸۹", // Extended Arabic-Indic Digits
		},



		{
			String: "Ⅰ Ⅱ Ⅲ Ⅳ Ⅴ Ⅵ Ⅶ Ⅷ Ⅸ Ⅹ Ⅺ Ⅻ Ⅼ Ⅽ Ⅾ Ⅿ",
		},
		{
			String: "ⅰ ⅱ ⅲ ⅳ ⅴ ⅵ ⅶ ⅷ ⅸ ⅹ ⅺ ⅻ ⅼ ⅽ ⅾ ⅿ",
		},
		{
			String: "ↀ ↁ ↂ Ↄ ↄ ↅ ↆ ↇ ↈ",
		},



		{
			String: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_.",
		},
		{
			String: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?",
		},
		{
			String: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_. !@#$%^&*()_+-={}[]|:;<>/?",
		},
	}


	for testNumber, test := range tests {

		parcel := ParcelFromString(test.String)


		if expected, actual := test.String, parcel.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}


		if expected, actual := test.String, string(parcel.Bytes()); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}


		if expected, actual := test.String, string(parcel.Runes()); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}


		actualReaderBytes, err := ioutil.ReadAll(parcel.Reader())
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v; for %q.", testNumber, err, test.String)
			continue
		}
		if expected, actual := test.String, string(actualReaderBytes); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}


		var buffer bytes.Buffer
		n, err := parcel.WriteTo(&buffer)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v; for %q.", testNumber, err, test.String)
			continue
		}
		if expected, actual := int64(len([]byte(test.String))), n; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for %q.", testNumber, expected, actual, test.String)
			continue
		}
		if expected, actual := test.String, buffer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q", testNumber, expected, actual)
			continue
		}


		var writer oitest.ShortWriter
		n, err = parcel.WriteTo(&writer)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v; for %q.", testNumber, err, test.String)
			continue
		}
		if expected, actual := int64(len([]byte(test.String))), n; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for %q.", testNumber, expected, actual, test.String)
			continue
		}
		if expected, actual := test.String, writer.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q", testNumber, expected, actual)
			continue
		}
	}
}


func TestStringParcelThen(t *testing.T) {


	tests := []struct{
		String string
		Thens []struct {
			Func func(Parcel)Parcel
			Expected string
		}
	}{
		{
			String: "AbCdE",
			Thens: []struct{
				Func func(Parcel)Parcel
				Expected string
			}{
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToUpper(parcel.String())
						return ParcelFromString(s)
					},
					Expected: "ABCDE",
				},
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToLower(parcel.String())
						return ParcelFromString(s)
					},
					Expected: "abcde",
				},
			},
		},



		{
			String: "Apple Banana Cherry",
			Thens: []struct{
				Func func(Parcel)Parcel
				Expected string
			}{
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToUpper(parcel.String())
						return ParcelFromString(s)
					},
					Expected: "APPLE BANANA CHERRY",
				},
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToLower(parcel.String())
						return ParcelFromString(s)
					},
					Expected: "apple banana cherry",
				},
			},
		},
	}


	for testNumber, test := range tests {

		parcel := ParcelFromString(test.String)

		for thenNumber, then := range test.Thens {

			if expected, actual := then.Expected, parcel.Then(then.Func).String(); expected != actual {
				t.Errorf("For test #%d and then #%d, expected %q, but actually got %q.", testNumber, thenNumber, expected, actual)
				continue
			}
		}
	}
}
