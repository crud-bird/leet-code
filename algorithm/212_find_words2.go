package algorithm

func FindWords2(board [][]byte, words []string) []string {
	trie := new(Trie)
	for _, w := range words {
		trie.put(w)
	}

	m := make([][]bool, len(board))
	for i := 0; i < len(m); i++ {
		m[i] = make([]bool, len(board[i]))
	}

	res := make(map[string]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			search(trie, board, i, j, make([]byte, 0, 8), m, &res)
		}
	}

	ss := make([]string, 0)
	for w := range res {
		ss = append(ss, w)
	}
	sortWords(ss)
	return ss
}

func search(t *Trie, board [][]byte, i, j int, prefix []byte, m [][]bool, res *map[string]bool) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || m[i][j] {
		return
	}

	prefix = append(prefix, board[i][j])
	node := t.searchPrefix(prefix)
	if node == nil {
		return
	}

	if node.IsEnd {
		(*res)[string(prefix)] = true
	}

	m[i][j] = true
	search(t, board, i+1, j, prefix, m, res)
	search(t, board, i-1, j, prefix, m, res)
	search(t, board, i, j+1, prefix, m, res)
	search(t, board, i, j-1, prefix, m, res)
	m[i][j] = false
}

//TrieNode ...
type TrieNode struct {
	Key      byte
	Children [26]*TrieNode
	IsEnd    bool
}

//Trie 定制化的trie
type Trie struct {
	root TrieNode
}

func (t *Trie) put(s string) {
	node := &t.root
	ss := []byte(s)
	for _, k := range ss {
		if node.Children[k-'a'] == nil {
			node.Children[k-'a'] = &TrieNode{Key: k}
		}
		node = node.Children[k-'a']
	}
	node.IsEnd = true
}

func (t *Trie) searchPrefix(ss []byte) *TrieNode {
	node := &t.root
	for _, s := range ss {
		if node.Children[s-'a'] == nil {
			return nil
		}
		node = node.Children[s-'a']
	}
	return node
}
