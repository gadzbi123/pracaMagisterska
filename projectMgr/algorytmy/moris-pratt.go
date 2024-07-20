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
				// sufix?
			}
		}
	}
	return
}

func moris_pratt(s []byte, substr []byte) (res []int) {
	lenS := len(s)
	lensubstr := len(substr)
	curr := -1
	preproc := make([]int, lensubstr+1)

	// dla wzorca obliczamy tablicę substrresubstrroc[]
	preproc[0] = -1
	for i := 1; i <= lensubstr; i++ {
		for (curr > -1) && (substr[curr] != substr[i-1]) {
			curr = preproc[curr]
		}
		curr++
		preproc[i] = curr
	}

	fmt.Println("preproc", preproc)
	// substroszukujemy substrozycji wzorca w łańcuchu
	curr = 0
	found := 0
	for i := 0; i < lenS; i++ {
		for (curr > -1) && (substr[curr] != s[i]) {
			curr = preproc[curr]
		}
		curr++
		if curr == lensubstr {
			for found < i-curr+1 {
				found++
			}
			res = append(res, found)
			found++
			curr = preproc[curr]
		}
	}
	return
}

func main() {
	// fmt.Println(moris_pratt([]byte("ALGORYTMTMA I STRUKTURY DANYCHYCH"), []byte("TMA")))
	fmt.Println(moris_pratt([]byte("BABAAABAABAAABABBABBBABBABABAABBAABAAABAABABABABAABBBAAAABBBBABBAABBBBBBABABAAA"), []byte("BBAAAABB")))
}
