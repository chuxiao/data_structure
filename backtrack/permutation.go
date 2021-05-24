package backtrack

func permutation(a []int) [][]int {
	res := make([][]int, 0)
	p := make([]int, 0)
	return permutation2(a, p, res)
}


func permutation2(a []int, p []int, res [][]int) [][]int {
	if len(p) == len(a) {
		p2 := make([]int, 0)
		p2 = append(p2, p...)
		res = append(res, p2)
		return res
	}
	for _, v := range a {
		if contain(p, v) {
			continue
		}
		p = append(p, v)
		res = permutation2(a, p, res)
		p = p[:len(p) - 1]
	}
	return res
}

func contain(p []int, v int) bool {
	for _, e := range p {
		if e == v {
			return true
		}
	}
	return false
}
