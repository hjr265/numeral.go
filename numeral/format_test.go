package numeral

import (
	"testing"
)

func TestFormatInt(t *testing.T) {
	tests := []struct{
		In Int
		Out string
	}{
		{
			Int{152, Fmt{Prec: 0, Suf: Byte}},
			"152B",
		},
		{
			Int{152000, Fmt{Prec: 0, Suf: Byte}},
			"152000B",
		},
		{
			Int{152, Fmt{Prec: 3, Suf: Abbr|Byte}},
			"152.000B",
		},
		{
			Int{152000, Fmt{Prec: 3, Suf: Abbr|Byte}},
			"148.438KB",
		},
		{
			Int{152000000, Fmt{Prec: 3, Suf: Abbr|Byte}},
			"144.958MB",
		},
		{
			Int{152, Fmt{Prec: 3, Suf: Abbr}},
			"152.000",
		},
		{
			Int{152000, Fmt{Prec: 3, Suf: Abbr}},
			"152.000k",
		},
		{
			Int{152000000, Fmt{Prec: 3, Suf: Abbr}},
			"152.000m",
		},
	}

	for _, test := range tests {
		s := test.In.String()
		if s != test.Out {
			t.Errorf("%#v.String() = %v; want %v", test.In, s, test.Out)
		}
	}
}
