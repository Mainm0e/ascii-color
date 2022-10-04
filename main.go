package main

import (
	"Color/errCheck"
	"Color/printletter"
	"os"
	"strings"

)

func main() {
	if len(os.Args) == 3 {
		arg3 := strings.Split(os.Args[2], "=")
		if arg3[0] == "--color" {
			if errCheck.Whatcolor(arg3[1]) == "err" {
				errCheck.UsageErr()
				os.Exit(0)
			}
			printletter.PrintArt(1, os.Args[1], errCheck.Checkcolor(errCheck.Whatcolor(arg3[1])))
		}

	} else if len(os.Args) == 4 {
		arg3 := strings.Split(os.Args[2], "=")
		if arg3[0] == "--color" && len(arg3) == 2 {
			if errCheck.Whatcolor(arg3[1]) == "err" {
				errCheck.UsageErr()
				os.Exit(0)
			}
			printletter.PrintArt(0,os.Args[1],errCheck.Checkcolor(errCheck.Whatcolor(arg3[1])))
		}
	} else {
		errCheck.UsageErr()
	}

}