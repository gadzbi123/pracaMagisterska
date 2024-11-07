package regular

import (
	"path/filepath"
	"testing"
)

func BenchmarkMorisPrattWindowWord(b *testing.B) {
	var founds = []string{}
	mp := &MorisPratt{}
	filepath.Walk(DIR, WalkAndFindByAlgo(mp, &founds, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expected", len(founds), 11598)
	}
}

func BenchmarkKurtMorisPrattWindowWord(b *testing.B) {
	var founds = []string{}
	kmp := &KurtMorisPratt{}
	filepath.Walk(DIR, WalkAndFindByAlgo(kmp, &founds, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expected", len(founds), 11598)
	}
}
func BenchmarkBoyerMooreWindowWord(b *testing.B) {
	var founds = []string{}
	bm := &BoyerMoore{}
	filepath.Walk(DIR, WalkAndFindByAlgo(bm, &founds, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expected", len(founds), 11598)
	}
}
