package numeral

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
	return ""
}