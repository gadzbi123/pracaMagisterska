package main

import (
	"fmt"
)

func preproces(s []byte) {

}
func boyer_moore(s []byte, substr []byte) (res []int) {
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
func main() {
	a := []byte("ALGORYTMY I STRUKTURY DANYCHYCH")
	b := []byte("YCH")
	fmt.Println(boyer_moore(a, b))
}
