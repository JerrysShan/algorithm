package structure

import "testing"

var bm = NewBitMap(1)

func TestSet(t *testing.T) {
	bm.Set(12)
	t.Log(bm.Test(12))
}

func TestClr(t *testing.T) {
	bm.Clr(12)
	t.Log(bm.Test(12))
}
