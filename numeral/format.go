package numeral

import (
	"strconv"
	"strings"
)

const (
	Abbr = 1 << iota
	Ord
	Byte
)

type Fmt struct {
	Prec int
	Suf int
}

type Int struct {
	int64
	Fmt Fmt
}

func (n Int) String() string {
	s := ""
	e := 0
	if n.Fmt.Suf&Abbr > 0 {
		b := 1e3
		if n.Fmt.Suf&Byte > 0 {
			b = 1024
		}

		v := float64(n.int64)
		for ; v >= b; e += 1 {
			v /= b
		}

		s += strconv.FormatFloat(v, 'f', n.Fmt.Prec, 64)

	} else {
		s += strconv.FormatInt(n.int64, 10)
		if n.Fmt.Prec > 0 {
			s += "." + strings.Repeat("0", n.Fmt.Prec)
		}
	}

	if n.Fmt.Suf & Byte > 0 {
		if n.Fmt.Suf & Abbr > 0 {
			switch e {
			case 1:
				s += "K"
			case 2:
				s += "M"
			case 3:
				s += "G"
			case 4:
				s += "T"
			case 5:
				s += "P"
			case 6:
				s += "E"
			case 7:
				s += "Z"
			case 8:
				s += "Y"
			}
		}
		s += "B"

	} else if n.Fmt.Suf & Abbr > 0 {
		switch e {
		case 1:
			s += "k"
		case 2:
			s += "m"
		case 3:
			s += "b"
		case 4:
			s += "t"
		}
	}

	return s
}