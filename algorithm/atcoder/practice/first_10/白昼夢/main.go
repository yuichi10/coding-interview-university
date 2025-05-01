package main

import (
	"fmt"
	"strings"
)

func main() {
	S := ""
	fmt.Scanf("%s", &S)

	// dreamerase
	// dreameraser
	// eraseraser
	//

	const DREAMER = "dreamer"
	const ERASER = "eraser"
	const DREAM = "dream"
	const ERASE = "erase"
	// s := strings.ReplaceAll(S, fmt.Sprintf("%s%s", DREAM, ERASER), "")
	// s = strings.ReplaceAll(s, fmt.Sprintf("%s%s", DREAM, ERASE), "")
	// s = strings.ReplaceAll(s, DREAMER, "")
	// s = strings.ReplaceAll(s, DREAM, "")
	// // s = strings.ReplaceAll(s, fmt.Sprintf("%s%s", ERASE, ERASE), "")
	// s = strings.ReplaceAll(s, ERASER, "")
	// s = strings.ReplaceAll(s, ERASE, "")
	s := strings.ReplaceAll(S, ERASER, "")
	s = strings.ReplaceAll(s, ERASE, "")
	s = strings.ReplaceAll(s, DREAMER, "")
	s = strings.ReplaceAll(s, DREAM, "")

	if s == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
