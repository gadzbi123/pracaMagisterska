package regular

import (
	"path/filepath"
	"testing"
)

func BenchmarkMorisPrattMainWord(b *testing.B) {
	var founds = []string{}
	mp := &MorisPratt{}
	filepath.Walk(DIR, WalkAndFindByAlgo(mp, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
func BenchmarkKurtMorisPrattMainWord(b *testing.B) {
	var founds = []string{}
	kmp := &KurtMorisPratt{}
	filepath.Walk(DIR, WalkAndFindByAlgo(kmp, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
func BenchmarkBoyerMooreMainWord(b *testing.B) {
	var founds = []string{}
	bm := &BoyerMoore{}
	filepath.Walk(DIR, WalkAndFindByAlgo(bm, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
