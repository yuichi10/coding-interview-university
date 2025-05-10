package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	S := ""
	fmt.Scanf("%s", &S)
	sliceS := strings.Split(S, "")
	length := len(S)
	max := 1 << (length - 1)
	sum := 0
	for i := 0; i < max; i++ {
		mSum := 0
		middle, _ := strconv.Atoi(sliceS[0])
		for j := 0; j < length-1; j++ {
			n, _ := strconv.Atoi(sliceS[j+1])
			if i>>j&1 == 0 {
				middle = middle*10 + n
			} else {
				mSum += middle
				middle = n
			}
		}
		if middle > 0 {
			mSum = mSum + middle
		}
		sum += mSum
	}
	fmt.Println(sum)
}
