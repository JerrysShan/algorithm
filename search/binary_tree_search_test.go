package search

import "testing"

var bst = &BinaryTreeSearch{}
var datas = []int{6, 1, 9, 2, 3, 4, 10, 15}

func TestBTSPut(t *testing.T) {
	for _, val := range datas {
		bst.Put(val)
	}
	bst.Show()
}

func TestBTSGet(t *testing.T) {
	node := bst.Get(4)
	if node != nil && node.Data == 4 {
		t.Log("get success")
	} else {
		t.Error("get data is wrong")
	}
}

func TestBTSSize(t *testing.T) {
	t.Log("test Binary_Tree_Search Size")
	if size := bst.Size(); size == len(datas) {
		t.Logf("return \" %d\"equal %d", size, len(datas))
	} else {
		t.Errorf("return\"%d\"not equal %d", size, len(datas))
	}
}

func TestBTSMin(t *testing.T) {
	t.Log("test Binary_Tree_Search Min should return 1")
	if data := bst.Min(); data == 1 {
		t.Logf("return \" %d\"", data)
	} else {
		t.Errorf("return\"%d\"", data)
	}
}

func TestBTSMax(t *testing.T) {
	t.Log("test Binary_Tree_Search Max should return 15")
	if data := bst.Max(); data == 15 {
		t.Logf("return \" %d\"", data)
	} else {
		t.Errorf("return\"%d\"", data)
	}
}

func TestBTSRank(t *testing.T) {
	t.Log("test Binary_Tree_Search Rank 15 should return 7")
	if data := bst.Rank(15); data == 7 {
		t.Logf("15 in bst rank should be 8th,return %d", data)
	} else {
		t.Errorf("15 in bst rank should be 7th,return %d", data)
	}
}

func TestBTSFloor(t *testing.T) {
	t.Log("test Binary_Tree_Search Floor")
	if node := bst.Floor(5); node.Data == 4 {
		t.Logf("low or equal 5 in bst should be 4,return result is %d ", node.Data)
	} else {
		t.Errorf("low or equal 5 in bst should be 4,but return result is %d", node.Data)
	}
}

func TestBTSCeiling(t *testing.T) {
	t.Log("test Binary_Tree_Search Ceiling")
	if node := bst.Ceiling(7); node.Data == 9 {
		t.Logf("greater or equal 7 in bst should b 9,return  is %d", node.Data)
	} else {
		t.Errorf("greater or equal 7 in bst should b 9,return  is %d", node.Data)
	}
}

func TestBTSIndex(t *testing.T) {
	t.Log("test Binary_Tree_Search Index")
	node, err := bst.Index(4)
	if err != nil {
		t.Error(err)
	}
	if node != nil && node.Data == 6 {
		t.Logf("rank 4th in bts should be 4,return %d", node.Data)
	} else {
		t.Errorf("rank 4th in bts should be 4,but return %d", node.Data)
	}
}

func TestBTSDeleteMin(t *testing.T) {
	t.Log("test Binary_Tree_Search DeleteMin")
	bst.DeleteMin()
	bst.Show()
}

func TestBTSDeleteMax(t *testing.T) {
	t.Log("test Binary_Tree_Search DeleteMax")
	bst.DeleteMax()
	bst.Show()
}

func TestBTSDelete(t *testing.T) {
	t.Log("test Binary_Tree_Search Delete")
	bst.Show()
}
