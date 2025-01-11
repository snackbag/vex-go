package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the world's most awesomest and best testing system!11!!")
	fmt.Println("Here's a list of all tests:")
	fmt.Println("- raylib_win")

	for {
		fmt.Print("\nEnter test (q): ")
		input, _ := reader.ReadString('\n')
		input = TrimSuffix(input, "\n")

		switch input {
		case "raylib_win":
			TestRaylibWin()
		default:
			fmt.Println("No test found by the name of '" + input + "'")
		}
	}
}

// TrimSuffix Thanks, stackoverflow
// https://stackoverflow.com/questions/8689425/how-to-remove-the-last-character-of-a-string-in-golang
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
