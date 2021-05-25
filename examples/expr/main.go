package main

import (
	"fmt"
	"strconv"
)

func isDigital(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func compare(op1, op2 rune) int {
	lv := func(op rune) int {
		switch op {
		case '+', '-':
			return 0
		case '*', '/':
			return 1
		default:
			panic("not supported.")
		}
	}
	comp := func(lv1, lv2 int) int {
		if lv1 > lv2 {
			return 1
		} else if lv1 == lv2 {
			return 0
		} else {
			return -1
		}
	}
	return comp(lv(op1), lv(op2))
}

func compute(op rune, n1, n2 int) int {
	switch op {
	case '+':
		return n1 + n2
	case '-':
		return n1 - n2
	case '*':
		return n1 * n2
	case '/':
		return n1 / n2
	default:
		panic("op not supported.")
	}
}
func main() {
	expr := "6*(5+(2+3)*8+3)"
	post := postfix(expr)
	fmt.Println(post)
	fmt.Println(computePostfix(post))
}

func postfix(expr string) string {
	output := []rune{}
	ops := NewOpStack(32)
	for _, r := range expr {
		if isDigital(r) {
			output = append(output, r)
		} else if r == '(' {
			ops.Push(r)
		} else if r == ')' {
			for {
				r2 := ops.Top()
				ops.Pop()
				if r2 == '(' {
					break
				}
				output = append(output, r2)
			}
		} else {
			op := string(r)
			_ = op
			if ops.IsEmpty() {
				ops.Push(r)
			} else {
				r2 := ops.Top()
				if r2 == '(' {
					ops.Push(r)
				} else if compare(r, r2) > 0 {
					ops.Push(r)
				} else {
					for r2 != '(' && compare(r, r2) <= 0 {
						ops.Pop()
						output = append(output, r2)
						r2 = ops.Top()
					}
					ops.Push(r)
				}
			}
		}
	}
	for !ops.IsEmpty() {
		r := ops.Top()
		ops.Pop()
		output = append(output, r)
	}
	return string(output)
}

func computePostfix(post string) int {
	ns := NewNumStack(32)
	for _, r := range post {
		if isDigital(r) {
			n, _ := strconv.Atoi(string(r))
			ns.Push(n)
		} else {
			n2 := ns.Top()
			ns.Pop()
			n1 := ns.Top()
			ns.Pop()
			val := compute(r, n1, n2)
			ns.Push(val)
		}
	}
	return ns.Top()
}
