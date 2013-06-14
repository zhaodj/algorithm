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

const (
	min myInt = myInt(500000)
	max myInt = myInt(700000)
)

func BenchmarkSearchSlice(b *testing.B) {
	rand.Seed(1)
	rp := rand.Perm(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res []myInt
		for _, v := range rp {
			if myInt(v).CompareTo(min) >= 0 && myInt(v).CompareTo(max) <= 0 {
				res = append(res, myInt(v))
			}
		}
	}
}

func BenchmarkSearchTree(b *testing.B) {
	var tree *Tree = &Tree{}
	rand.Seed(1)
	rp := rand.Perm(1000000)
	for _, v := range rp {
		tree.Add(myInt(v))
	}
	//fmt.Println(tree.Level())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.SearchRange(min, max)
	}
}
