package fx

import (
	"math"
	"math/rand"
	"testing"
)

func TestCordicSinCos(t *testing.T) {
	N := 1024
	for i := 0; i < N; i++ {
		fi := rand.Float64()
		phi := Float64(fi)
		x, y := cordicSinCos(phi)
		se := math.Abs(x.Float64() - math.Sin(fi))
		ce := math.Abs(y.Float64() - math.Cos(fi))
		if se > 1e-14 {
			t.Errorf("sin %s cosf %f err %0.12f\n", x, math.Sin(fi), se)
		}
		if ce > 1e-14 {
			t.Errorf("cos %s cosf %f err %0.12f\n", x, math.Cos(fi), ce)
		}
	}
}

func TestCordicAtan(t *testing.T) {
	N := 32
	for i := 0; i < N; i++ {
		fi := -rand.Float64()
		fat := math.Atan(fi)
		phi := Float64(fi)
		at := cordicAtan(phi, One)
		t.Logf("fat %f at %s\n", fat, at)
	}
}
