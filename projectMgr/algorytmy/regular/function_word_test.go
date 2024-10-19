package regular

import (
	"path/filepath"
	"testing"
)

func BenchmarkMorisPrattFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var founds = []string{}
		filepath.Walk(DIR, WalkAndFindByAlgoAndWord(moris_pratt, &founds, []byte("function")))
		if len(founds) != 32619 {
			b.Fatal("result did not match with expeceted", len(founds), 32619)
		}
	}
}
