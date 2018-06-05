package fx

import (
	"fmt"
	"math"
	"strings"
)

const (
	iBits  = 63 - FrBits
	frMask = (1 << FrBits) - 1
	iMask  = ((1 << iBits) - 1) << FrBits
)

const (
	Sign = -1
	Zero = 0
	One  = 1 << FrBits
	Iota = 1
	Max  = iMask | frMask
)

type T int64

func Int(i int) T {
	i &= (1 << iBits) - 1
	return T(i) << FrBits
}

func Float64(f float64) T {
	s := T(Iota)
	if f < 0.0 {
		f = -f
		s = Sign
	}
	fone := float64(One)
	g := math.Floor(fone*f + 0.5)
	return s * T(int64(g))
}

func (t T) Float64() float64 {
	fone := float64(One)
	return float64(t) / fone
}

func (t T) String() string {
	s := ""
	if t < 0 {
		t = -t
		s = "-"
	}
	ds := decimal(int64(t & frMask))
	f := string(ds[1:])
	i := fmt.Sprintf("%d", (t>>FrBits)+T((ds[0]-'0')))
	return strings.Join([]string{s, i, ".", f}, "")
}

func decimal(v int64) []byte {
	u := v
	var buf [32]byte
	buf[0] = '0'
	n := 1
	o := int64(One)
	o1 := o
	for {
		o1 = o / 10
		if o1 <= u {
			break
		}
		buf[n] = '0'
		n++
		o = o1
	}
	u *= 10
	for o > 0 {
		buf[n] = '0' + byte(u/o)
		m := n
		for buf[m] == '0'+10 {
			buf[m] = '0'
			buf[m-1]++
			m--
		}
		u = u % o
		o /= 10
		n++
	}
	m := n - 1
	for m >= 1 && buf[m] == '0' {
		m--
	}
	n = m + 1
	if n == 1 {
		n = 2
	}
	return buf[:n]
}

// Frac returns the fractional part of a.
func (a T) Frac() T {
	return a & (frMask | Sign)
}

// Int returns the int part of a.
func (a T) Int() T {
	return a & (iMask | Sign)
}

// Round returns a rounded to the nearest integer
// as an int64.
func (a T) Round() int64 {
	c := T(Zero)
	if a&(One>>1) != 0 {
		c = Iota
	}
	return int64((a >> FrBits) + c)
}

// Abs returns the absolute value of a.
func (a T) Abs() T {
	if a < 0 {
		a = -a
	}
	return a
}

// Mul computes multiplication of a*b.
func (a T) Mul(b T) T {
	s, aa, ab := mulSign(a, b)
	low, high := mulBits(uint64(aa), uint64(ab))
	f := (high << (iBits + 1)) | (low >> FrBits) + ((low >> (FrBits - 1)) & 1)
	return s * T(f)
}

// Div computes division.  Note if b is a power of 2,
// then you can just shift.
func (a T) Div(b T) T {
	s, aa, bb := mulSign(a, b)
	u := &u128{hi: uint64(aa) >> (64 - FrBits), lo: uint64(aa) & ((1 << (iBits + 1)) - 1)}
	u.lo <<= FrBits
	v := u.divBits(uint64(bb))
	return s * T(v)
}

// Inv computes and returns 1/a.
func (a T) Inv() T {
	return T(One).Div(a)
}

// Sqrt computes the square root of a.  If a is negative
// Sqrt panics.
func Sqrt(a T) T {
	if a < 0 {
		panic("sqrt negative")
	}
	const steps = 6

	var x T = One
	var t T
	for i := 0; i <= steps; i++ {
		t = a.Div(x) + x
		x = t >> 1
	}
	return x
}

// Sin computes sin(a).
//
// a must be in the range
//
//  [-2*Pi..2*Pi]
//
// or Sin panics
func Sin(a T) T {
	s, _ := SinCos(a)
	return s
}

// Cos computes cos(a)
//
// a must be in the range
//
//  [-2*Pi..2*Pi]
//
// or Cos panics
func Cos(a T) T {
	_, c := SinCos(a)
	return c
}

// SinCos computes sin(a), cos(a)
//
// a must be in the range
//
//  [-2*Pi..2*Pi]
//
// or SinCos panics
func SinCos(a T) (T, T) {
	return cordicSinCos(a)
}

// Tan computes the tangent of a.
//
// a must be in the range
//
//  [-2*Pi..2*Pi]
//
// or Tan panics
func Tan(a T) T {
	s, c := SinCos(a)
	return s.Div(c)
}

// Atan2 computes the arctangent of x/y
//
// Atan2 distinguishes the result based on the signs of
// both x and y.
func Atan2(x, y T) T {
	return cordicAtan2(x, y)
}

// Atan computes the arc-tangent of a,
// ie the angle w for which sin(w)/cos(w) = a.
func Atan(a T) T {
	return Atan2(a, One)
}

// returns sign, abs(a), abs(b) as T.
func mulSign(a, b T) (s, absA, absB T) {
	s = T(1)
	if a^b < 0 {
		s = -1
	}
	if a < 0 {
		a = -a
	}
	absA = a
	if b < 0 {
		b = -b
	}
	absB = b
	return
}

// Returns the full 128 bit result of x * y in low, high.
//
// Note(wsc) some people do something similar and quote hackers delight, i didn't grok
// the casts.  This is just derived from multiplying 32 bit chunks piecewise
// and then factoring away the computations specific to low and replacing them
// with x*y.  Apparently, about the same number of operations, and the same
// number of multiplies as the Hacker's Delight version -- without casts!.
func mulBits(x, y uint64) (low, high uint64) {
	const (
		shift = 32
		mask  = (1 << shift) - 1
	)
	low = x * y

	xl, xh := x&mask, x>>shift
	yl, yh := y&mask, y>>shift

	ll := xl * yl
	lh := xl * yh
	hl := xh * yl
	hh := xh * yh

	t := ll>>shift + lh&mask + hl&mask

	c := t >> shift
	t = lh>>shift + hl>>shift + c + hh&mask
	high = t & mask
	c = t >> shift
	high |= (hh>>shift + c) << shift
	return
}
