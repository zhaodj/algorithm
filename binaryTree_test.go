package algorithm

import (
	//"fmt"
	"log"
	"math/rand"
	"testing"
)

type myInt int

func (t myInt) CompareTo(v Comparable) int {
	value, ok := v.(myInt)
	if !ok {
		log.Fatal("type assert error")
	}
	return (int)(t - value)
}

func TestTree(t *testing.T) {
	var tree *Tree = &Tree{}
	rand.Seed(1)
	rp := rand.Perm(10)
	t.Log(rp)
	for _, v := range rp {
		tree.Add(myInt(v))
	}
	if !tree.Contain(myInt(2)) {
		t.Error("2 not found")
	}
	t.Log(tree.SearchRange(myInt(2), myInt(5)))
	tree.Print()
}

func TestAVLTree(t *testing.T) {
	var tree *AVLTree = &AVLTree{}
	rand.Seed(1)
	rp := rand.Perm(10)
	for _, v := range rp {
		tree.Add(myInt(v))
	}
	tree.Print()
}

const (
	TEST_SIZE int   = 1000000
	TEST_MIN  myInt = myInt(500000)
	TEST_MAX  myInt = myInt(700000)
)

func sliceContain(s []int, v Comparable) bool {
	for _, item := range s {
		if myInt(item).CompareTo(v) == 0 {
			return true
		}
	}
	return false
}

func newBinaryTree() *Tree {
	var tree *Tree = &Tree{}
	rand.Seed(1)
	rp := rand.Perm(TEST_SIZE)
	for _, v := range rp {
		tree.Add(myInt(v))
	}
	return tree
}

func newAVLTree() *AVLTree {
	var tree *AVLTree = &AVLTree{}
	rand.Seed(1)
	rp := rand.Perm(TEST_SIZE)
	for _, v := range rp {
		tree.Add(myInt(v))
	}
	return tree
}

func BenchmarkSearchSlice(b *testing.B) {
	rand.Seed(1)
	rp := rand.Perm(TEST_SIZE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res []myInt
		for _, v := range rp {
			if myInt(v).CompareTo(TEST_MIN) >= 0 && myInt(v).CompareTo(TEST_MAX) <= 0 {
				res = append(res, myInt(v))
			}
		}
	}
}

func BenchmarkSearchTree(b *testing.B) {
	tree := newBinaryTree()
	//fmt.Println(tree.Level())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.SearchRange(TEST_MIN, TEST_MAX)
	}
}

func BenchmarkSearchAVLTree(b *testing.B) {
	tree := newAVLTree()
	//fmt.Println(tree.Level())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.SearchRange(TEST_MIN, TEST_MAX)
	}
}

func BenchmarkSliceContain(b *testing.B) {
	rand.Seed(1)
	rp := rand.Perm(TEST_SIZE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceContain(rp, TEST_MAX)
	}
}

func BenchmarkBTreeContain(b *testing.B) {
	tree := newBinaryTree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Contain(TEST_MAX)
	}
}

func BenchmarkAVLTreeContain(b *testing.B) {
	tree := newAVLTree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Contain(TEST_MAX)
	}
}
