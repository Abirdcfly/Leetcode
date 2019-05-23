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


