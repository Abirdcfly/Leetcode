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
