package regular

import (
	"path/filepath"
	"testing"
)

func BenchmarkMorisPrattMainWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(moris_pratt, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
func BenchmarkKurtMorisPrattMainWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(kurt_moris_pratt, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
func BenchmarkBoyerMooreMainWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(boyer_moore, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
