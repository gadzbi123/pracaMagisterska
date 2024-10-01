package regular

import "fmt"

func kurt_moris_pratt(s []byte, substr []byte) (res []int) {
	substrlen := len(substr)
	slen := len(s)
	KMPNext := make([]int, substrlen+1)
	KMPNext[0] = -1
	b := -1
	for i := 1; i <= substrlen; i++ {
		for (b > -1) && (substr[b] != substr[i-1]) {
			b = KMPNext[b]
		}
		b++
		// Addition check here
		if (i == substrlen) || (substr[i] != substr[b]) {
			KMPNext[i] = b
		} else {
			KMPNext[i] = KMPNext[b]
		}
	}

	pp := 0
	b = 0
	for i := 0; i < slen; i++ {
		for (b > -1) && (substr[b] != s[i]) {
			b = KMPNext[b]
		}
		b++
		if b == substrlen {
			for pp < i-b+1 {
				pp++
			}
			res = append(res, pp)
			pp++
			b = KMPNext[b]
		}
	}
	return
}

func main2() {
	fmt.Println(kurt_moris_pratt(
		[]byte("BABAAABAABAAABABBABBBABBABABAABBAABAAABAABABABABAABBBAAAABBBBABBAABBBBBBABABAAA"),
		[]byte("BBAAAABB"),
	))
}
