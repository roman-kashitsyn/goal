package bitset

import (
	"testing"
	"testing/quick"
)

func qc(t *testing.T, f interface{}) {
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInitialization(t *testing.T) {
	qc(t, func(bits []bool) bool {
		bs := ToBits(bits)
		for i, b := range bits {
			if b != bs.IsSet(i) {
				return false
			}
		}
		return true
	})
}

func TestReset(t *testing.T) {
	qc(t, func(bits []bool) bool {
		bs := ToBits(bits)
		bs.Reset()
		for i := 0; i < bs.Len(); i++ {
			if bs.IsSet(i) {
				return false
			}
		}
		return true
	})
}

func TestLogicalNot(t *testing.T) {
	qc(t, func(bits []bool) bool {
		s := ToBits(bits)
		s.Not()
		for i, b := range bits {
			if b == s.IsSet(i) {
				return false
			}
		}
		return true
	})
}

func TestLogicalAndOr(t *testing.T) {
	qc(t, func(lhs, rhs []bool) bool {
		l, r := ToBits(lhs), ToBits(rhs)
		andRes, orRes := l.Copy(), l.Copy()

		andRes.And(r)
		orRes.Or(r)

		n := min(len(lhs), len(rhs))
		for i := 0; i < n; i++ {
			if andRes.IsSet(i) != (lhs[i] && rhs[i]) {
				return false
			}
			if orRes.IsSet(i) != (lhs[i] || rhs[i]) {
				return false
			}
		}
		return true
	})
}

func TestToBoolConversion(t *testing.T) {
	qc(t, func(bools []bool) bool {
		s := ToBits(bools)
		r := s.ToBools()
		if len(bools) != len(r) {
			return false
		}
		for i, b := range bools {
			if b != r[i] {
				return false
			}
		}
		return true
	})
}
