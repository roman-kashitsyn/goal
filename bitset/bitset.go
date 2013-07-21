// Package bitset provides operations on bit sets.
package bitset

const byteLen = 32

// BitSet is a fixed-size array of bits.
type BitSet struct {
	data   []uint32
	length int
}

// New creates a new bit set of given length.
func New(length int) *BitSet {
	n := (length + byteLen - 1) / byteLen
	data := make([]uint32, n)
	return &BitSet{data, length}
}

// Copy creates a copy of a bit set.
func (s *BitSet) Copy() *BitSet {
	data := make([]uint32, len(s.data))
	copy(data, s.data)
	return &BitSet{data, s.length}
}

// Set sets given bit to 1.
func (s *BitSet) Set(i int) {
	n := index(i, s.length)
	offset := uint(i % byteLen)
	s.data[n] = s.data[n] | 1<<offset
}

// Unset sets given bit to 0.
func (s *BitSet) Unset(i int) {
	n := index(i, s.length)
	offset := uint(i % byteLen)
	s.data[n] &= ^(1 << offset)
}

// IsSet checks whether given bit is 1.
func (s *BitSet) IsSet(i int) bool {
	n := index(i, s.length)
	offset := uint(i % byteLen)
	return s.data[n]&(1<<offset) != 0
}

// Reset sets all the bits to 0.
func (s *BitSet) Reset() *BitSet {
	for i, _ := range s.data {
		s.data[i] = 0
	}
	return s
}

// Not inverts all the bits in a set.
func (s *BitSet) Not() {
	for i := 0; i < len(s.data); i++ {
		s.data[i] = ^s.data[i]
	}
}

// And sets all bits of receiving bit set into logical and (&)
// operation of receiving and recieved sets.
func (lhs *BitSet) And(rhs *BitSet) {
	n := min(len(lhs.data), len(rhs.data))
	for i := 0; i < n; i++ {
		lhs.data[i] &= rhs.data[i]
	}
}

// Or sets all bits of receiving bit set into logical or (|) operation
// of receiving and recieved sets.
func (lhs *BitSet) Or(rhs *BitSet) {
	n := min(len(lhs.data), len(rhs.data))
	for i := 0; i < n; i++ {
		lhs.data[i] |= rhs.data[i]
	}
}

// ToBits creates a new bit set from a given bool slice.
func ToBits(bits []bool) *BitSet {
	bs := New(len(bits))
	for i, b := range bits {
		if b {
			bs.Set(i)
		}
	}
	return bs
}

// ToBools converts a bit set into a slice of bools.
func (s *BitSet) ToBools() []bool {
	bits := make([]bool, s.length)
	for i := 0; i < s.length; i++ {
		bits[i] = s.IsSet(i)
	}
	return bits
}

// Len returns length of a bit set.
func (s *BitSet) Len() int {
	return s.length
}

func index(i, n int) int {
	if i < 0 || n <= i {
		panic("Set index out of bounds")
	}
	return i / byteLen
}

func min(x, y int) int {
	m := x
	if (y < m) {
		m = y
	}
	return m
}