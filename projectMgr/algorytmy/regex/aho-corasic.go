package regex

type Trie struct {
	curr   string
	next   []*Trie
	output bool
}

type Bytes = []byte

func FillTrie(subchains []string) (root *Trie) {
	if len(subchains) == 0 {
		return
	}
	var prev *Trie = root
	var temp *Trie = root
	for _, sc := range subchains {
		for i, ch := range sc {
			if prev == nil {
				root = &Trie{curr: ""}
			} else {
				temp = &Trie{curr: prev.curr + string(ch)}
				prev.next = append(prev.next, temp)
				if i == len(sc)-1 {
					temp.output = true
				}
			}

		}

	}

	return
}

func main() {
	var łańcuch []byte = []byte("abedget")
	var podłańcuchy [][]byte = [][]byte{
		[]byte("ab"),
		[]byte("about"),
		[]byte("at"),
		[]byte("ate"),
		[]byte("be"),
		[]byte("bed"),
		[]byte("edge"),
		[]byte("get"),
	}
}
