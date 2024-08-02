package v3

import "testing"

func TestBitset(t *testing.T) {
	var b BitSet
	var exp []bool

	b = BitSet{
		Buf: []byte{0xFF}, // 1111_1111
		Len: 8,
	}
	exp = []bool{true, true, true, true, true, true, true, true}
	for i := 0; i < 8; i++ {
		if b.Get(i) != exp[i] {
			t.Errorf("BitSet.Get(%d) should return %v", i, exp[i])
		}
	}
	if b.Get(9) {
		t.Errorf("BitSet.Get(9) should return false")
	}

	b = BitSet{
		Buf: []byte{0xA3}, // 1010_0011
		Len: 8,
	}
	exp = []bool{true, true, false, false, false, true, false, true}
	for i := 0; i < 8; i++ {
		if b.Get(i) != exp[i] {
			t.Errorf("BitSet.Get(%d) should return %v", i, exp[i])
		}
	}
	if b.Get(9) {
		t.Errorf("BitSet.Get(9) should return false")
	}

	b = BitSet{
		Buf: []byte{0x00}, // 0000_0000
		Len: 8,
	}
	exp = []bool{false, false, false, false, false, false, false, false}
	for i := 0; i < 8; i++ {
		if b.Get(i) != exp[i] {
			t.Errorf("BitSet.Get(%d) should return %v", i, exp[i])
		}
	}
	if b.Get(9) {
		t.Errorf("BitSet.Get(9) should return false")
	}
}
