package fx

import (
	"math"
	"math/rand"
	"testing"
)

func TestCordicSinCos(t *testing.T) {
	N := 1024
	for i := 0; i < N; i++ {
		fi := (rand.Float64() - 0.5) * 2 * math.Pi
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

func TestCordicAtan2YNeg(t *testing.T) {
	N := 512
	for i := 0; i < N; i++ {
		fi := (rand.Float64() - 0.5) * 2 * math.Pi
		fat := math.Atan2(fi, -1.0)
		phi := Float64(fi)
		at := cordicAtan2(phi, -One)
		if math.Abs(at.Float64()-fat) > 1e-14 {
			t.Errorf("%s fat %f at %s\n", phi, fat, at)
		}
	}
}

func TestCordicAtan2YPos(t *testing.T) {
	N := 512
	for i := 0; i < N; i++ {
		fi := (rand.Float64() - 0.5) * 2 * math.Pi
		fat := math.Atan2(fi, 1.0)
		phi := Float64(fi)
		at := cordicAtan2(phi, One)
		if math.Abs(at.Float64()-fat) > 1e-14 {
			t.Errorf("%s fat %f at %s\n", phi, fat, at)
		}
	}
}
