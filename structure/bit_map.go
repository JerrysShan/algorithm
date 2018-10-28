package structure

const SHIFT = 5
const MASK uint32 = 32

type BitMap struct {
	arr []uint32
}

func NewBitMap(length int) *BitMap {
	return &BitMap{
		arr: make([]uint32, length),
	}
}

func (b *BitMap) Set(i uint32) {
	index := i >> SHIFT // i/32
	b.arr[index] = b.arr[index] | 1<<(i&MASK)
}

func (b *BitMap) Clr(i uint32) {
	index := i >> SHIFT
	b.arr[index] &= b.arr[index] & ^(1 << (i & MASK))
}

func (b *BitMap) Test(i uint32) uint32 {
	index := i >> SHIFT
	return (b.arr[index] & (1 << (i & MASK)))
}
