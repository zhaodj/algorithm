package algorithm

import "testing"

func TestSearch(t *testing.T) {
	sl := make([]int, 20)
	for i := 4; i < len(sl)+4; i++ {
		sl[i-4] = i
	}
	if Search(sl, 4) != 0 {
		t.Error("index start failed")
	} else if Search(sl, 23) != 19 {
		t.Error("index last failed")
	} else if Search(sl, 20) != 16 {
		t.Error("index 16 failed")
	} else if Search(sl, 21) != 17 {
		t.Error("index 17 failed")
	} else if Search(sl, 29) != -1 {
		t.Error("index not exists failed")
	}
}

func BenchmarkSearch(b *testing.B) {
	sl := make([]int, 2000000)
	for i := 0; i < len(sl); i++ {
		sl[i] = i
	}
	for j := 0; j < b.N; j++ {
		Search(sl, j)
	}
}
