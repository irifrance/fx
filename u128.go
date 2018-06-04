package fx

import (
	"fmt"
	"math/bits"
)

type u128 struct {
	lo uint64
	hi uint64
}

func (n *u128) divBits(m uint64) uint64 {
	if n.hi == 0 {
		return n.lo / m
	}
	nlh := bits.Len64(n.hi)
	ml := bits.Len64(m)
	m128 := &u128{lo: m}
	ns := 64 + (nlh - ml)
	m128.lshift(uint(ns))

	q128 := &u128{}
	for ns >= 0 {
		q128.lshift(1)
		if m128.leq(n) {
			q128.lo |= 1
			n.sub(m128)
		}
		m128.rshift(1)
		ns--
	}
	return q128.lo //
}

func (u *u128) less(v *u128) bool {
	if u.hi < v.hi {
		return true
	}
	if u.hi > v.hi {
		return false
	}
	return u.lo < v.lo
}

func (u *u128) leq(v *u128) bool {
	if u.hi == v.hi && u.lo == v.lo {
		return true
	}
	return u.less(v)
}

func (u *u128) add(v *u128) *u128 {
	return u
}

func (u *u128) sub(v *u128) *u128 {
	u.lo -= v.lo
	c := (((u.lo & v.lo) & 1) + (v.lo >> 1) + u.lo>>1) >> 63
	u.hi -= v.hi + c
	return u
}

func (u *u128) rshift(s uint) *u128 {
	if s < 64 {
		u.lo >>= s
		u.lo |= u.hi << (64 - s)
		u.hi >>= s
		return u
	}
	u.lo = u.hi >> (s - 64)
	u.hi = 0
	return u
}

func (u *u128) lshift(s uint) *u128 {
	if s < 64 {
		u.hi <<= s
		u.hi |= u.lo >> (64 - s)
		u.lo <<= s
		return u
	}
	u.hi = u.lo << (s - 64)
	u.lo = 0
	return u
}

func (u *u128) String() string {
	return fmt.Sprintf("%064b%064b", u.hi, u.lo)
}
