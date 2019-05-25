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
