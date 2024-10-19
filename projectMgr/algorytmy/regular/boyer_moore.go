package main

import (
	"bytes"
	"fmt"
)

func boyer_moore_old(s []byte, substr []byte) (res []int) {
	slen := len(s)
	substrlen := len(substr)
	preproces(substr)
	for i := 0; i < slen; {
		for j := 0; j < substrlen; j++ {
			fmt.Println("start", s[i+j], substr[j], i+j, j)
			if i+j >= slen {
				fmt.Println("Max len return")
				return
			}
			if s[i+j] == substr[j] {
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

// Ref: https://github.com/cubicdaiya/bms/blob/master/bms.go
func preproces(substr []byte) map[byte]int {
	l := len(substr)
	table := make(map[byte]int)

	for i := 0; i < l-1; i++ {
		j := substr[i]
		table[j] = l - i - 1
	}

	return table
}

func boyer_moore(str, substr []byte) []int {
	table := preproces(substr)
	i := 0
	len_str := len(str)
	len_substr := len(substr)
	results := []int{}

	if len_str == 0 || len_substr == 0 || len_str < len_substr {
		return results
	}

	if bytes.Equal(str, substr) {
		return []int{0}
	}

loop:
	for i+len_substr <= len_str {
		for j := len_substr - 1; j >= 0; j-- {
			if str[i+j] != substr[j] {
				if _, ok := table[str[i+j]]; !ok {
					if j == len_substr-1 {
						i += len_substr
					} else {
						i += len_substr - j - 1
					}
				} else {
					n := table[str[i+j]] - (len_substr - j - 1)
					if n <= 0 {
						i++
					} else {
						i += n
					}
				}
				goto loop
			}
		}

		results = append(results, i)
		if _, ok := table[str[i+len_substr-1]]; ok {
			i += table[str[i+len_substr-1]]
		} else {
			i += len_substr
		}
	}

	return results
}
func main() {
	a := []byte("ALGORYTMY I STRUKTURY DANYCHYCH")
	b := []byte("YCH")
	fmt.Println(boyer_moore(a, b))
}
