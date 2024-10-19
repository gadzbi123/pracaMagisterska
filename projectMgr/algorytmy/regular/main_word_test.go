package regular

import (
	"path/filepath"
	"testing"
)

func BenchmarkMorisPrattMainWord(b *testing.B) {
	var founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(moris_pratt, &founds, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expeceted", len(founds), 19716)
	}

}
