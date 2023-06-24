package main

import (
	"asciiart"
	"fmt"
	"os"
)

func main() {
	// Check if input is correct
	var rawInput string
	sourceFile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	ColorChosen := ""
	indexofcolorwords := -1
	lengthofcolorwords := -1
	outputErr := "\nUsage: go run . [OPTION] [STRING]\n\nExample: go run . --color=red <letters to be colored> something \n\n"
	if len(os.Args[1]) <= 7 {
		fmt.Print(outputErr)
			return
	}
	if len(os.Args) == 2 { // [STRING]
		rawInput = os.Args[1]
	} else if len(os.Args) == 3 { // [STRING] [BANNER]
		rawInput = os.Args[2]
		if os.Args[1][:8] == "--color=" {
			ColorChosen = os.Args[1]
		} else {
			fmt.Print(outputErr)
			return
		}
		indexofcolorwords-- // if indexofcolorwords is -2 the whole word well be colored
	} else if len(os.Args) == 4 { // [OPTION] [STRING]
		rawInput = os.Args[3]
		if os.Args[1][:8] == "--color=" {
			ColorChosen = os.Args[1]
		} else {
			fmt.Print(outputErr)
			return
		}
	} else { // ERROR (FOR OUTPUT USE)
		fmt.Print(outputErr)
		return
	}
	if len(os.Args) == 4 {
		letterstobecolored := os.Args[2]
		indexofcolorwords = asciiart.Index(rawInput, letterstobecolored)
		lengthofcolorwords = len(os.Args[2]) - 1
	}
	// Main function: Splitting (split string based on newline position)
	// ∟--> Sub function: Formatting (change input to allow use of newline & qoutation marks)
	splitInput := asciiart.LineSplitter(rawInput, asciiart.InputFormatter)

	// Main function: Printing (printing the row of characters within input string)
	// ∟--> Sub function: Parsing (parsing the data of the 8 rows to print sequentially)
	asciiart.RowPrinter(splitInput, ColorChosen, indexofcolorwords, lengthofcolorwords, sourceFile, asciiart.RowParser)
}
