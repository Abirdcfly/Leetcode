import "strings"

func simplifyPath(path string) string {
	path1 := strings.Split(path, "/")
	npath := make([]string, 0)
	for _, i := range path1 {
		if i == "" || i == "." {
			continue
		} else if i == ".." {
			if len(npath) != 0 {
				npath = npath[:len(npath)-1]
			}
		} else {
			npath = append(npath, i)
		}
	}
	return "/" + strings.Join(npath, "/")
}
