import "math"

func findAnagrams(s string, p string) []int {
	f := func(s string) (res int) {
		/* time out
		   ns := strings.Split(s,"")
		   sort.Strings(ns)
		   return strings.Join(ns,"")
		*/
		for _, i := range s {
			res += int(math.Pow(10, float64(i-'a')))
		}
		//fmt.Println(s,res)
		return
	}
	np := f(p)
	r := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if i+len(p) <= len(s) && f(s[i:i+len(p)]) == np {
			r = append(r, i)
		}
	}
	return r
}

//滑动窗口法https://www.youtube.com/watch?v=86fQQ7rVGxA
func findAnagrams(s string, p string) []int {
	lp := len(p)
	mp := [26]int{}
	ms := [26]int{}
	res := make([]int, 0)
	for _, i := range p {
		mp[i-'a']++
	}
	for index, i := range s {
		if index >= lp {
			ms[s[index-lp]-'a']--
		}
		ms[i-'a']++
		if mp == ms {
			res = append(res, index-lp+1)
		}
	}
	return res
}


