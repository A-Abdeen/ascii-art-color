package asciiart

import "fmt"

func RowPrinter(splitInput []string, i int, sourceFile []byte, f func([]byte, int) []byte) {
	fullRowData := ""
	for _, singleLine := range splitInput { // to print one line at a time
		if singleLine != "" {
				for _, charRune := range singleLine { // to combine the prespective line of each char to the next
					rowLocation := (((int(charRune) - 32) * 9) + i)
					charRowData := f(sourceFile, rowLocation)
					fullRowData = fullRowData + string(charRowData)
				}
				return fullRowData
				fullRowData = ""
			}
		} else {
			fmt.Print("\n")
		}
	}
