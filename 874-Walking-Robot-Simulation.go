func robotSim(commands []int, obstacles [][]int) int {
	i, j := 0, 0
	d := 0 //0^ 1> 2下 3<
	ri, rj, flag := 0, 0, false
	max := 0
	for _, x := range commands {
		if x >= 1 && x <= 99 {
			if d == 0 {
				if len(obstacles) == 0 {
					j += x
					if max < i*i+j*j {
						max = i*i + j*j
					}
					continue
				}
				t1 := j + x
				for _, t := range obstacles {
					oi, oj := t[0], t[1]
					if ri, rj, flag = inpath(i, j, i, j+x, oi, oj); flag {
						if t1 > rj {
							t1 = rj
						}
					}
				}
				i, j = i, t1
			} else if d == 1 {
				if len(obstacles) == 0 {
					i += x
					if max < i*i+j*j {
						max = i*i + j*j
					}
					continue
				}
				t1 := i + x
				for _, t := range obstacles {
					oi, oj := t[0], t[1]
					if ri, rj, flag = inpath(i, j, i+x, j, oi, oj); flag {
						if t1 > ri {
							t1 = ri
						}
					}
				}
				i, j = t1, j
			} else if d == 2 {
				if len(obstacles) == 0 {
					j -= x
					if max < i*i+j*j {
						max = i*i + j*j
					}
					continue
				}
				t1 := j - x
				for _, t := range obstacles {
					oi, oj := t[0], t[1]
					if ri, rj, flag = inpath(i, j, i, j-x, oi, oj); flag {
						if t1 < rj {
							t1 = rj
						}
					}
				}
				i, j = i, t1
			} else {
				if len(obstacles) == 0 {
					i -= x
					if max < i*i+j*j {
						max = i*i + j*j
					}
					continue
				}
				t1 := i - x
				for _, t := range obstacles {
					oi, oj := t[0], t[1]
					if ri, rj, flag = inpath(i, j, i-x, j, oi, oj); flag {
						if t1 < ri {
							t1 = ri
						}
					}
				}
				i, j = t1, j
			}
		} else if x == -1 {
			d = (d + 1) % 4
			if d < 0 {
				d += 4
			}
		} else {
			d = (d - 1) % 4
			if d < 0 {
				d += 4
			}
		}
		if max < i*i+j*j {
			max = i*i + j*j
		}
	}
	return max
}

func inpath(si, sj, ei, ej, oi, oj int) (ri, rj int, flag bool) {
	if si == ei && si == oi && sj < oj && oj <= ej {
		ri, rj = ei, oj-1
		flag = true
	} else if si == ei && si == oi && ej <= oj && oj < sj {
		ri, rj = ei, oj+1
		flag = true
	} else if sj == ej && sj == oj && si < oi && oi <= ei {
		ri, rj = oi-1, ej
		flag = true
	} else if sj == ej && sj == oj && ei <= oi && oi < si {
		ri, rj = oi+1, ej
		flag = true
	} else {
		ri, rj = ei, ej
	}
	return
}

//更好理解的自己写的版本
func robotSim(commands []int, obstacles [][]int) int {
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	d := 0
	x, y := 0, 0
	max := 0
	for _, i := range commands {
		if i == -1 {
			d = (d + 1) % 4
		} else if i == -2 {
			d = (d - 1 + 4) % 4
		} else {
			tx, ty := dx[d]*1, dy[d]*1
			for j := 0; j < i; j++ {
				stop := false
				for _, item := range obstacles {
					if x+tx == item[0] && y+ty == item[1] {
						stop = true
					}
				}
				if !stop {
					x, y = x+tx, y+ty
				} else {
					break
				}
			}
			maxt := x*x + y*y
			if maxt > max {
				max = maxt
			}
		}
	}
	return max
}

//用map会快很多
func robotSim(commands []int, obstacles [][]int) int {
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	d := 0
	x, y := 0, 0
	max := 0
	obs := make(map[int]bool)
	for _, v := range obstacles {
		obs[v[0]*600000+v[1]] = true // 因为题目中限定了x,y的范围所以可以这么写
	}
	for _, i := range commands {
		if i == -1 {
			d = (d + 1) % 4
		} else if i == -2 {
			d = (d - 1 + 4) % 4
		} else {
			tx, ty := dx[d]*1, dy[d]*1
			for j := 0; j < i; j++ {
				stop := false
				if _, e := obs[(x+tx)*600000+(y+ty)]; e {
					stop = true
				}
				if !stop {
					x, y = x+tx, y+ty
				} else {
					break
				}
			}
			maxt := x*x + y*y
			if maxt > max {
				max = maxt
			}
		}
	}
	return max
}
