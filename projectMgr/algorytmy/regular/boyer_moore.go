package regular

import (
	"fmt"
	"slices"

	"github.com/gadzbi123/pracaMagisterska/algorytmy/utils"
)

/*
	func (bm *BoyerMoore) Find(s []byte, substr []byte) (res []int) {
		slen := len(s)
		substrlen := len(substr)
		if bm.preproc == nil {
			table := make(map[byte]int)

			for i := 0; i < substrlen-1; i++ {
				j := substr[i]
				table[j] = substrlen - i - 1
			}
			bm.preproc = table

		}
		for i := 0; i < slen; {
			for j := substrlen - 1; j >= 0; j-- {
				// fmt.Println("start", s[i+j], substr[j], i+j, j)
				if s[i+j] != substr[j] {
					fmt.Println("continue", s[i+j], substr[j], i+j, j)
					if j == substrlen-1 {
						res = append(res, i)
						fmt.Println("found", i)
						i += substrlen
					}
				} else if j == 0 {
					i++
					goto SKIP_SUBSTR
				} else {
					i += substrlen
				}
			}
		SKIP_SUBSTR:
		}
		return
	}
*/

type BoyerMoore struct {
	preproc []int
}

var _ utils.AlgoStruct = &BoyerMoore{}

// Ref: https://github.com/cubicdaiya/bms/blob/master/bms.go
func (bm *BoyerMoore) Find(str, substr []byte) (res []int) {
	i := 0
	len_str := len(str)
	len_substr := len(substr)
	if bm.preproc == nil {
		l := len(substr)
		table := make([]int, 256)

		for i := 0; i < l-1; i++ {
			j := substr[i]
			table[j] = l - i - 1
		}
		bm.preproc = table
	}
loop:
	for i+len_substr <= len_str {
		for j := len_substr - 1; j >= 0; j-- {
			if str[i+j] != substr[j] {
				if loc := bm.preproc[str[i+j]]; loc == 0 {
					if j == len_substr-1 {
						i += len_substr
					} else {
						i += len_substr - j - 1
					}
				} else {
					n := loc - len_substr + j + 1
					if n <= 0 {
						i++
					} else {
						i += n
					}
				}
				goto loop
			}
		}
		res = append(res, i)
		if v := bm.preproc[str[i+len_substr-1]]; v != 0 {
			i += v
		} else {
			i += len_substr
		}
	}
	return
}
func main3() {
	bm := &BoyerMoore{}
	a := []byte("ALGORYTMY CH I STRUKTURY DANYCHYCHCH CHCHCH")
	res1 := (bm.Find(a, []byte("CH")))
	exp1 := []int{10, 29, 32, 34, 37, 39, 41}
	if !slices.Equal(res1, exp1) {
		panic(fmt.Sprintf("res1 not equal: %v != %v", res1, exp1))
	}
}
