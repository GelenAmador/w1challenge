package mario

import (
	"fmt"
	"strings"
)

func RunMarioPyramid() {
	var height int

	for {
		fmt.Println("Pyramid height:")
		fmt.Scanf("%d", &height)

		if height >= 1 && height <= 8 {
			break
		}
		fmt.Println("Please enter a number between 1 and 8")
	}

	for i := 1; i <= height; i++ {
		fmt.Print(strings.Repeat(" ", height-i))
		fmt.Println(strings.Repeat("#", i))
	}
}
