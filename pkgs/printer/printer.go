package printer

import "fmt"

var (
	colorReset = "\033[0m"

    colorRed = "\033[31m"
    colorGreen = "\033[32m"
    colorYellow = "\033[33m"
    colorBlue = "\033[34m"
    colorPurple = "\033[35m"
    colorCyan = "\033[36m"
    colorWhite = "\033[37m"
)



func PrintGreen(value string) {
	fmt.Println(colorGreen+value+colorReset)
}


func PrintRed(value string) {
	fmt.Println(colorRed+value+colorReset)
}


func PrintYellow(value string) {
	fmt.Println(colorYellow+value+colorReset)
}

func PrintBlue(value string) {
	fmt.Println(colorBlue+value+colorReset)
}

func PrintPurple(value string) {
	fmt.Println(colorPurple+value+colorReset)
}

func PrintCyan(value string) {
	fmt.Println(colorCyan+value+colorReset)
}

func PrintWhite(value string) {
	fmt.Println(colorWhite+value+colorReset)
}