package errCheck

import "fmt"

func UsageErr() {
	fmt.Println("Usage: go run . [STRING] [OPTION]")
	fmt.Println()
	fmt.Println("EX: go run . something --color=<color>")

}

func Checkcolor(c string) string {
	switch c {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "purple":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	default:
		return "\033[0m"
	}
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
	case "Purple", "purple", "PURPLE":
		return "purple"
	case "Cyan", "cyan", "CYAN":
		return "cyan"
	case "White", "white", "WHITE":
		return "white"
	default:
		return "err"
	}
}
