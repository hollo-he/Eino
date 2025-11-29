package utils

import "fmt"

// 柔和的蓝紫渐变
var softColors = []int{
	69, 75, 81, 87, 93, 99, 105, 111,
}

func PrintBanner(text string) {
	for i, r := range text {
		color := softColors[i%len(softColors)]
		fmt.Printf("\033[38;5;%dm%s\033[0m", color, string(r))
	}
	fmt.Println()
}
