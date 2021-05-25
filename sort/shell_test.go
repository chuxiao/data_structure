package sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	a := []int{64, 8, 216, 512, 27, 729, 0, 1, 343, 125}
	ShellSort(a)
	fmt.Println(a)
}
