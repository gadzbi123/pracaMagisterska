package main

import (
	"bytes"
	"fmt"
)

func prefix_sufix(s []byte) (i int) {
	defer fmt.Printf("string:%q,sufix:%v\n", s, i)
	n := len(s)
	for i = n - 1; i >= 0; i-- {
		if bytes.Compare(s[0:i], s[n-i:n]) == 0 {
			return i
		}
	}
	return
}
func moris_pratt_slow(s []byte, substr []byte) (res []int) {
	slen := len(s)
	substrlen := len(substr)
	for i := 0; i < slen; i++ {
		for j := 0; j < substrlen; j++ {
			if slen <= i+j {
				// fmt.Println("too long")
				return
			}
			// fmt.Println("comparing", s[i+j], substr[j], i, j)
			if s[i+j] != substr[j] {
				if j > 0 {
					pattern_part := s[i : i+j+1]
					bb := prefix_sufix(pattern_part)
					mov_len := len(pattern_part) - bb - 1
					// fmt.Println("movlen", mov_len, len(pattern_part), bb)
					i += mov_len
				}
				break
			}
			if j == substrlen-1 {
				res = append(res, i)
				//sufix?
			}
		}
	}
	return
}
func main() {
	// fmt.Println(moris_pratt([]byte("ALGORYTMTMA I STRUKTURY DANYCHYCH"), []byte("TMA")))
	fmt.Println(moris_pratt([]byte("BABAAABAABAAABABBABBBABBABABAABBAABAAABAABABABABAABBBAAAABBBBABBAABBBBBBABABAAA"), []byte("BBAAA")))
}
