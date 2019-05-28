type Trie struct {
	d           [26]*Trie
	isEndofWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
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

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
