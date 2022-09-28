package errCheck

import "fmt"

func UsageErr() {
	fmt.Println("Usage: go run . [STRING] [OPTION]")
	fmt.Println()
	fmt.Println("EX: go run . something --color=<color>")

}

func Whatcolor(c string) string {
	switch c {
	case "Black", "black", "BLACK":
		return "black"
	case "Red", "red", "RED":
		return "red"
	case "Green", "green", "GREEN":
		return "green"
	case "Yellow", "yellow", "YELLOW":
		return "yellow"
	case "Blue", "blue", "BLUE":
		return "blue"
	case "Magenta", "magenta", "MAGENTA":
		return "magenta"
	case "Cyan", "cyan", "CYAN":
		return "cyan"
	case "White", "white", "WHITE":
		return "white"
	default:
		return "err"
	}
}
