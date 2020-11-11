package algo

import "fmt"

type Fraction struct {
	Numerator   int
	Denumerator int
}

var result string

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denumerator)
}

func Gen(l Fraction, r Fraction, num int) {
	var f Fraction
	f = Fraction{l.Numerator + r.Numerator, l.Denumerator + r.Denumerator}
	if f.Denumerator <= num {
		Gen(l, f, num)
		fmt.Printf("%s ", f)
		Gen(f, r, num)
	}
}
