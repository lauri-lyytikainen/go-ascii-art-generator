package draw

import "fmt"

func Draw(width int, height int) {
	for _ = range height {
		for _ = range width {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
