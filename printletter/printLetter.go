package printletter

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)

func FineSpecial(){
	letter := strings.Split(os.Args[3], ",")
	fmt.Println(len(letter))

}


func PrintArt(n int,s,c string){
	fmt.Print(Printer(n,s,c))
}

func Printer(n int, s, c string)string {
		x := c
	file, err := os.Open("standard.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	if n != 1{
		x = "\033[0m"
	}
	var l1, l2, l3, l4, l5, l6, l7, l8 string

	Special := []int{4,6,7,3,1}
	sp := 0
	for i := 0; i < len(s); i++ {
		var y int
		
		if i == Special[sp] {
			x = c //terminal that what we want
			sp++
		} else {
			x = "\033[0m"
		}
		

		letter := s[i]
		DecN := int(letter - 32)
		for i := 0; i < DecN; i++ {
			y = y + 9
		}

		stringNumb := y + 1

		for i := 1; i <= 8; i++ {
			line := i
			for numb, eachline := range txtlines {
				if numb == stringNumb {
					switch line {
					case 1:
						l1 = l1 + string(x) + eachline + string("\033[0m")
					case 2:
						l2 = l2 + string(x) + eachline + string("\033[0m")
					case 3:
						l3 = l3 + string(x) + eachline + string("\033[0m")
					case 4:
						l4 = l4 + string(x) + eachline + string("\033[0m")
					case 5:
						l5 = l5 + string(x) + eachline + string("\033[0m")
					case 6:
						l6 = l6 + string(x) + eachline + string("\033[0m")
					case 7:
						l7 = l7 + string(x) + eachline + string("\033[0m")
					case 8:
						l8 = l8 + string(x) + eachline + string("\033[0m") 
					default:
					}

				}

			}
			stringNumb++
		}
	}




	asciiArt := l1 + "\n" + l2 + "\n" + l3 + "\n" + l4 + "\n" + l5 + "\n" + l6 + "\n" + l7 + "\n" + l8 + "\n"  
	return asciiArt
}


