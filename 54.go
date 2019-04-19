//54. Spiral Matrix
func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	var r, c, di int //(r,c)目前坐标，认定目前坐标未入结果，di为方向坐标，0，1，2，3按逆时针转动，dr，dc为当方向转dao时，横纵坐标变化。，比如，di=0，dr[0]=0.dc[0]=1指的是方向为3点钟方向，r不变，c+1
	dr := [...]int{0, 1, 0, -1}
	dc := [...]int{1, 0, -1, 0}
	rlen := len(matrix)
	clen := len(matrix[0])
	path := make([]bool, rlen*clen)
	for i := 0; i < rlen*clen; i++ {
		res = append(res, matrix[r][c])
		path[r*clen+c] = true
		nr, nc := r+dr[di], c+dc[di]
		if 0 <= nr && nr < rlen && 0 <= nc && nc < clen && path[nr*clen+nc] == false {
			r, c = nr, nc
		} else {
			di = (di + 1) % 4
			r, c = r+dr[di], c+dc[di]
		}
	}
	return res
}
