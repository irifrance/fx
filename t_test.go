package fx

import (
	"math/rand"
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
	N := 1024
	for i := 0; i < N; i++ {
		n := T(rand.Int63n((1 << 63) - 1))

		m := n.Mul(One)
		if n != m {
			t.Errorf("%d: One isn't identify for %s\no %b\nm %b", i, n, n, m)
		}
	}
}

func TestDivId(t *testing.T) {
	N := 1024
	for i := 0; i < N; i++ {
		n := T(rand.Int63n((1 << 63) - 1))
		m := n.Div(One)
		if n != m {
			t.Errorf("%d: One isn't identify for %s\no %b\nm %b", i, n, n, m)
		}
	}
}

func TestInvPo2(t *testing.T) {
	for i := uint(0); i < iBits; i++ {
		n := T(One << i)
		m := n.Inv().Inv()
		if m != n {
			t.Errorf("%d: inv(inv(%s)) gave %s\n", i, n, m)
		}
	}
}

func TestInvMulClose(t *testing.T) {
	for i := 1; i < (1 << iBits); i++ {
		n := Int(i)
		f := n.Mul(n.Inv())
		e := One - f
		if e > Iota<<9 {
			t.Logf("%s * %s = %s\n", n, n.Inv(), e)
		}

	}
}
