package regular

import (
	"path/filepath"
	"testing"
)

func BenchmarkMorisPrattWindowWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(moris_pratt, &founds, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expeceted", len(founds), 11598)
	}
}

func BenchmarkKurtMorisPrattWindowWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(kurt_moris_pratt, &founds, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expeceted", len(founds), 11598)
	}
}
func BenchmarkBoyerMooreWindowWord(b *testing.B) {
	b.SkipNow()
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(boyer_moore, &founds, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expeceted", len(founds), 11598)
	}
}
