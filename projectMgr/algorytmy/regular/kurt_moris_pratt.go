package regular

import "fmt"

type KurtMorisPratt struct {
	preproc []int
}

var _ AlgoStruct = &KurtMorisPratt{}

func (kmp *KurtMorisPratt) Find(s []byte, substr []byte) (res []int) {
	substrlen := len(substr)
	slen := len(s)
	if kmp.preproc == nil {
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
		kmp.preproc = KMPNext
	}

	pp := 0
	b := 0
	for i := 0; i < slen; i++ {
		for (b > -1) && (substr[b] != s[i]) {
			b = kmp.preproc[b]
		}
		b++
		if b == substrlen {
			for pp < i-b+1 {
				pp++
			}
			res = append(res, pp)
			pp++
			b = kmp.preproc[b]
		}
	}
	return
}

func main2() {
	kmp := KurtMorisPratt{}
	fmt.Println(kmp.Find(
		[]byte("BABAAABAABAAABABBABBBABBABABAABBAABAAABAABABABABAABBBAAAABBBBABBAABBBBBBABABAAA"),
		[]byte("BBAAAABB"),
	))
}
