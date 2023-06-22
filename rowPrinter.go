package asciiart

import "fmt"

func RowPrinter(splitInput []string, ColorChosen string, indexofcolorwords int, lengthofcoloredwords int, sourceFile []byte, f func([]byte, int) []byte) {
	fullRowData := ""
	fmt.Println(indexofcolorwords)
	fmt.Println(lengthofcoloredwords)
	for _, singleLine := range splitInput { // to print one line at a time
		if singleLine != "" {
			for i := 2; i < 10; i++ {
				for x, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
					if (x >= indexofcolorwords && x <= (indexofcolorwords + lengthofcoloredwords)){
						fmt.Print(fullRowData)
						Color(string(charRowData), ColorChosen, lengthofcoloredwords)
						charRowData = []byte{}
						fullRowData = ""
					}
					fullRowData = fullRowData + string(charRowData)
				}
				if indexofcolorwords == -2 {
					Color(fullRowData, ColorChosen, indexofcolorwords)
				} else {
				fmt.Println(fullRowData)
			}
				fullRowData = ""
			}
		} else {
			fmt.Print("\n")
		}
	}
}
