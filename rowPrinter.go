package asciiart

import "fmt"

func RowPrinter(splitInput []string, ColorChosen string, indexOfColorWords int, lengthOfColorWords int, sourceFile []byte, f func([]byte, int) []byte) {
	fullRowData := ""
	for _, singleLine := range splitInput { // to print one line at a time
		if singleLine != "" {
			for i := 2; i < 10; i++ {
				for x, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
					if (indexOfColorWords != -1) && (x >= indexOfColorWords && x <= (indexOfColorWords+lengthOfColorWords)) {
						fmt.Print(fullRowData)
						Color(string(charRowData), ColorChosen, lengthOfColorWords)
						charRowData = []byte{}
						fullRowData = ""
					}
					fullRowData = fullRowData + string(charRowData)
				}
				if indexOfColorWords == -2 {
					Color(fullRowData, ColorChosen, indexOfColorWords)
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
