package regular

import (
	"path/filepath"
	"testing"

	"github.com/gadzbi123/pracaMagisterska/algorytmy/utils"
)

func BenchmarkMorisPrattWindowWord(b *testing.B) {
	var founds = []string{}
	mp := &MorisPratt{}
	filepath.Walk(utils.PERF_DIR, utils.WalkAndFindByAlgo(mp, &founds, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expected", len(founds), 11598)
	}
}

func BenchmarkKurtMorisPrattWindowWord(b *testing.B) {
	var founds = []string{}
	kmp := &KurtMorisPratt{}
	filepath.Walk(utils.PERF_DIR, utils.WalkAndFindByAlgo(kmp, &founds, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expected", len(founds), 11598)
	}
}
func BenchmarkBoyerMooreWindowWord(b *testing.B) {
	var founds = []string{}
	bm := &BoyerMoore{}
	filepath.Walk(utils.PERF_DIR, utils.WalkAndFindByAlgo(bm, &founds, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expected", len(founds), 11598)
	}
}
