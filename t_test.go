package fx

import (
	"testing"
)

func TestString(t *testing.T) {
	t.Logf("%s", T(One))
	t.Logf("%s", T(Iota))
	t.Logf("%s", T(One>>1))
	t.Logf("%s", T(One>>2))
	t.Logf("%s", T(One+One+One).Mul(One>>1))
}

func TestMulId(t *testing.T) {
	N := 1
	for i := 0; i < N; i++ {
		//n := T(rand.Int63n((1 << 63) - 1))
		n := T(One)

		m := n.Mul(One)
		if n != m {
			t.Errorf("%d: One isn't identify for %s\no %b\nm %b", i, n, n, m)
		}
	}
}

func TestDivId(t *testing.T) {
	t.Logf("%s", T(One).Div(One))
}
