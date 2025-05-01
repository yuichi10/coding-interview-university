package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	N := 0
	fmt.Scanf("%d", &N)
	sc := bufio.NewScanner(os.Stdin)
	cx := 0
	cy := 0
	ct := 0
	no := false
	for i := 0; i < N; i++ {
		sc.Scan()
		nums := strings.Split(sc.Text(), " ")
		t, _ := strconv.Atoi(nums[0])
		x, _ := strconv.Atoi(nums[1])
		y, _ := strconv.Atoi(nums[2])
		difft := t - ct
		diffx := x - cx
		diffy := y - cy
		if !no {
			if (diffx+diffy)%2 == 0 && difft%2 == 0 && int(math.Abs(float64(diffx+diffy))) <= difft {
				cx = x
				cy = y
				ct = t
			} else if math.Abs(float64((diffx+diffy)%2)) == 1 && difft%2 == 1 && int(math.Abs(float64(diffx+diffy))) <= difft {
				cx = x
				cy = y
				ct = t
			} else {
				no = true
			}
		}
	}
	if no {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
