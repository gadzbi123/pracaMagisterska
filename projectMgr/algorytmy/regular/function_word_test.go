package regular

import (
	"path/filepath"
	"testing"
)

func BenchmarkMorisPrattFunctionWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgo(moris_pratt, &founds, []byte("function")))
	if len(founds) != 32619 {
		b.Fatal("result did not match with expected", len(founds), 32619)
	}
}

func BenchmarkKurtMorisPrattFunctionWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgo(kurt_moris_pratt, &founds, []byte("function")))
	if len(founds) != 32619 {
		b.Fatal("result did not match with expected", len(founds), 32619)
	}
}
func BenchmarkBoyerMooreFunctionWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgo(boyer_moore, &founds, []byte("function")))
	if len(founds) != 32619 {
		b.Fatal("result did not match with expected", len(founds), 32619)
	}
}
