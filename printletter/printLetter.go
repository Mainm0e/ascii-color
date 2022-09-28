package printletter

import (
	"fmt"
	"os"
	"strings"
)

func Specifiled(c string) {
	splitSpecifiled := strings.Split(os.Args[3], ",")

	letterColor := c

	for i := 0; i < len(splitSpecifiled); i++ {
		fineLetter := strings.Split(os.Args[1], splitSpecifiled[i])
		for j := 0; j < len(fineLetter); j++ {
			Printletter(fineLetter[j], "none")
			Printletter(splitSpecifiled[i], letterColor)
		}
	}
}

func Printletter(s, c string) {
	Color := c
	String := s

	fmt.Println("this is string", String)
	fmt.Println("this is color", Color)
}
