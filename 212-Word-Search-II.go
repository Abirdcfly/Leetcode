
func findWords(board [][]byte, words []string) []string {
	t := ConstructorTrie()
	for _, word := range words {
		t.Insert(word)
	}
	res := make([]string, 0)
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			search(i, j, len(board), len(board[0]), board, visited, &t, &[]byte{}, &res)
		}
	}
	return res
}

func search(r, c, lr, lc int, board [][]byte, visited [][]bool, t *Trie, cur *[]byte, res *[]string) {
	if r >= lr || c >= lc || r < 0 || c < 0 {
		return
	}

	if visited[r][c] {
		return
	} else {
		visited[r][c] = true
	}

	*cur = append(*cur, board[r][c])
	if scur := string(*cur); (*t).StartsWith(scur) {
		if (*t).Search(scur) {
			exist := false
			for _, i := range *res {
				if i == scur {
					exist = true
					break
				}
			}
			if !exist {
				*res = append(*res, scur)
			}
		}
		search(r, c-1, lr, lc, board, visited, t, cur, res)
		search(r-1, c, lr, lc, board, visited, t, cur, res)
		search(r, c+1, lr, lc, board, visited, t, cur, res)
		search(r+1, c, lr, lc, board, visited, t, cur, res)
	}
	*cur = (*cur)[:len(*cur)-1]
	visited[r][c] = false
	return
}

type Trie struct {
	d           [26]*Trie
	isEndofWord bool
}

/** Initialize your data structure here. */
func ConstructorTrie() Trie {
	return Trie{
		d: [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, i := range word {
		x := this.d[i-'a']
		if x == nil {
			this.d[i-'a'] = &Trie{}
			this = this.d[i-'a']
		} else {
			this = x
		}
	}
	this.isEndofWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, i := range word {
		x := this.d[i-'a']
		if x == nil {
			return false
		}
		this = x
	}
	if !this.isEndofWord {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, i := range prefix {
		x := this.d[i-'a']
		if x == nil {
			return false
		}
		this = x
	}
	return true
}

