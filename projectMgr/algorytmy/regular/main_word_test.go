package regular

import (
	"path/filepath"
	"testing"

	"github.com/gadzbi123/pracaMagisterska/algorytmy/utils"
)

func BenchmarkMorisPrattMainWord(b *testing.B) {
	var founds = []string{}
	mp := &MorisPratt{}
	filepath.Walk(utils.PERF_DIR, utils.WalkAndFindByAlgo(mp, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
func BenchmarkKurtMorisPrattMainWord(b *testing.B) {
	var founds = []string{}
	kmp := &KurtMorisPratt{}
	filepath.Walk(utils.PERF_DIR, utils.WalkAndFindByAlgo(kmp, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
func BenchmarkBoyerMooreMainWord(b *testing.B) {
	var founds = []string{}
	bm := &BoyerMoore{}
	filepath.Walk(utils.PERF_DIR, utils.WalkAndFindByAlgo(bm, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
