package parcels


import (
	"io/ioutil"
	"strings"

	"testing"
)


func TestFromString(t *testing.T) {

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
	}


	for testNumber, test := range tests {

		parcel := FromString(test.String)


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
			t.Errorf("For test #%d, did not expect an error, but actually got one: %v", testNumber, err)
			continue
		}
		if expected, actual := test.String, string(actualReaderBytes); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}


func TestStringParcelMap(t *testing.T) {


	tests := []struct{
		String string
		Mappings []struct {
			Func func(Parcel)Parcel
			Expected string
		}
	}{
		{
			String: "AbCdE",
			Mappings: []struct{
				Func func(Parcel)Parcel
				Expected string
			}{
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToUpper(parcel.String())
						return FromString(s)
					},
					Expected: "ABCDE",
				},
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToLower(parcel.String())
						return FromString(s)
					},
					Expected: "abcde",
				},
			},
		},



		{
			String: "Apple Banana Cherry",
			Mappings: []struct{
				Func func(Parcel)Parcel
				Expected string
			}{
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToUpper(parcel.String())
						return FromString(s)
					},
					Expected: "APPLE BANANA CHERRY",
				},
				{
					Func: func(parcel Parcel) Parcel {
						s := strings.ToLower(parcel.String())
						return FromString(s)
					},
					Expected: "apple banana cherry",
				},
			},
		},
	}


	for testNumber, test := range tests {

		parcel := FromString(test.String)

		for mappingNumber, mapping := range test.Mappings {

			if expected, actual := mapping.Expected, mapping.Func(parcel).String(); expected != actual {
				t.Errorf("For test #%d and mapping #%d, expected %q, but actually got %q.", testNumber, mappingNumber, expected, actual)
				continue
			}
		}
	}
}
