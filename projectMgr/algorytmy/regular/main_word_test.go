package regular

import (
	"fmt"
	"path/filepath"
	"testing"
)

func BenchmarkMorisPrattMainWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgo(moris_pratt, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
func BenchmarkKurtMorisPrattMainWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgo(kurt_moris_pratt, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
}
func BenchmarkBoyerMooreMainWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgo(boyer_moore, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expected", len(founds), 19716)
	}
	// PrintFounds(founds)
	fmt.Println("LOL")
	b.SetParallelism(8)
}
