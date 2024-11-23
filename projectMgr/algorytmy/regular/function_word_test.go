package regular

import (
	"path/filepath"
	"testing"

	"github.com/gadzbi123/pracaMagisterska/algorytmy/utils"
)

func BenchmarkMorisPrattFunctionWord(b *testing.B) {
	var founds = []string{}
	mp := &MorisPratt{}
	filepath.Walk(utils.PERF_DIR, utils.WalkAndFindByAlgo(mp, &founds, []byte("function")))
	if len(founds) != 32619 {
		b.Fatal("result did not match with expected", len(founds), 32619)
	}
}

func BenchmarkKurtMorisPrattFunctionWord(b *testing.B) {
	var founds = []string{}
	kmp := &KurtMorisPratt{}
	filepath.Walk(utils.PERF_DIR, utils.WalkAndFindByAlgo(kmp, &founds, []byte("function")))
	if len(founds) != 32619 {
		b.Fatal("result did not match with expected", len(founds), 32619)
	}
}
func BenchmarkBoyerMooreFunctionWord(b *testing.B) {
	var founds = []string{}
	bm := &BoyerMoore{}
	filepath.Walk(utils.PERF_DIR, utils.WalkAndFindByAlgo(bm, &founds, []byte("function")))
	if len(founds) != 32619 {
		b.Fatal("result did not match with expected", len(founds), 32619)
	}
}
