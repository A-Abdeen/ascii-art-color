package asciiart

import "fmt"

func RowPrinter(splitInput []string, ColorChosen string, indexofcolorwords int, lengthofcolorwords int, sourceFile []byte, f func([]byte, int) []byte) {
	fullRowData := ""
	for _, singleLine := range splitInput { // to print one line at a time
		if singleLine != "" {
			for i := 2; i < 10; i++ {
				for x, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
					if (indexofcolorwords != -1 ) && (x >= indexofcolorwords && x <= (indexofcolorwords + lengthofcolorwords)){
						fmt.Print(fullRowData)
						Color(string(charRowData), ColorChosen, lengthofcolorwords)
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
