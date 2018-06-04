package fx

// NB: see qXX.go for cordic constant tables "cordicAtans"
// and "cordicKs"

// w: angle in radians
func cordicSinCos(w T) (sin, cos T) {
	if w < -Pi/2 || w > Pi/2 {
		panic("w out of range")
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
	return y.Mul(cordicKs[FrBits-1]), x.Mul(cordicKs[FrBits-1])
}

func cordicAtan(x, y T) T {
	var tmp, z T
	for i := uint(0); i < FrBits; i++ {
		if y < 0 {
			tmp = x - y>>i
			y = y + x>>i
			x = tmp
			z -= cordicAtans[i]
		} else {
			tmp = x + y>>i
			y = y - x>>i
			x = tmp
			z += cordicAtans[i]
		}
	}
	return z
}
