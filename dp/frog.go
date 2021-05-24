package dp

func frog(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	n1, n2 := 1, 2
	for i := 3; i <= n; i++ {
		k := n2
		n2 += n1
		n1 = k
	}
	return n2
}
