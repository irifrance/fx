package fx

import (
	"fmt"
	"math"
	"testing"
)

func TestCordic(t *testing.T) {
	phi := T(One << 8).Inv()
	x, y := cordicSinCos(phi)
	fmt.Printf("sin %s cos %s sinf %f cosf %f\n", x, y, math.Sin(phi.Float64()), math.Cos(phi.Float64()))
}
