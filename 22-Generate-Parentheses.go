func generateParenthesis(n int) []string {
	if n < 0 {
		return nil
	}
	res := make([]string, 0)
	t := make(map[string]bool)
	t["()"] = true
	res = append(res, "()")
	if n == 1 {
		return res
	}
	for i := 1; i < n; i++ {
		lenr := len(res)
		for _, j := range res {
			for k := 0; k < len(j)+1; k++ {
				n := j[0:k] + "()" + j[k:]
				if _, e := t[n]; !e {
					res = append(res, n)
					t[n] = true
				}
			}
			delete(t, j)
		}
		res = res[lenr:]
	}
	return res
}

func generateParenthesis(n int) []string {
	r := make([]string, 0)
	help(&r, "", 0, 0, n)
	return r
}

// https://www.youtube.com/watch?v=XF0wh8M2A6E
func help(res *[]string, s string, l, r, n int) {
	if len(s) == 2*n {
		*res = append(*res, s)
		return
	}
	if l < n {
		help(res, s+"(", l+1, r, n)
	}
	if r < l {
		help(res, s+")", l, r+1, n)
	}
}

//https://zxi.mytechroad.com/blog/searching/leetcode-22-generate-parentheses/
//核心实际都是对生成的那棵树，剪掉不合适的枝叶，后的结果。前面的代码类似BFS，上面一行链接是DFS
