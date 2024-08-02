package v3

type BitSet struct {
	Buf []byte
	Len int
}

func (s *BitSet) Get(idx int) bool {
	if idx > s.Len {
		return false
	}
	byteIdx := idx >> 3
	bitIdx := idx & 7
	return (s.Buf[byteIdx]>>bitIdx)&1 != 0
}
