// Copyright 2018 Iri France SAS. All rights reserved.  Use of this source code
// is governed by a license that can be found in the License file.

package fx

// implementations of
// CORDIC (Coordinate rotation digital computer) algorithm
// for fixed point computations of trig functions.

// NB: see qXX.go for cordic constant tables "cordicAtans"
// and "cordicKs" and cmd/genq/ for their generation.

// w: angle in radians
func cordicSinCos(w T) (sin, cos T) {
	if w < -Pi || w > Pi {
		// what to do about this?  need to support
		// something like modf I guess...
		panic("outside range")
	}
	s := T(Iota)
	if w > Pi/2 {
		w -= Pi
		s = -s
	} else if w < -Pi/2 {
		w += Pi
		s = -s
	}
	x := T(One)
	y := T(Zero)

	var tmp, z T
	for i := uint(0); i < FrBits; i++ {
		if z < w {
			tmp = x - y>>i
			y = y + x>>i
			x = tmp
			z += cordicAtans[i]
		} else {
			tmp = x + y>>i
			y = y - x>>i
			x = tmp
			z -= cordicAtans[i]
		}
	}
	// more precise to multiply here since the constant
	// is less than one
	return s * y.Mul(cordicKs[FrBits-1]), s * x.Mul(cordicKs[FrBits-1])
}

func cordicAtan2(x, y T) T {
	var z = T(Pi / 2) // start at top of range to scale w.r.t. pi w/out Mul
	if y == 0 {
		if x > 0 {
			return z
		}
		if x < 0 {
			return -z
		}
		panic("undefined atan")
	}
	s := T(Iota)
	var add = T(Zero)
	if y < 0 {
		y = -y
		s = -s
		add = -Pi
		if x < 0 {
		}
	}
	if x < 0 {
		s = -s
		x = -x
	}
	var tmp T
	for i := uint(0); i < FrBits; i++ {
		if y < 0 {
			tmp = x - y>>i
			y = y + x>>i
			x = tmp
			z += cordicAtans[i]
		} else {
			tmp = x + y>>i
			y = y - x>>i
			x = tmp
			z -= cordicAtans[i]
		}
	}
	return (s * (z + add))
}
