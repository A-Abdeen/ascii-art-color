package asciiart

import (
    "fmt"
)

func Color(charRowData string, ColorChosen string, indexofcolorwords int){
	color := "\033[0m"
    if ColorChosen == "--color=black"{
        color = "\033[30m"}
	if ColorChosen == "--color=red"{
    color = "\033[31m"}
    if ColorChosen == "--color=green"{
    color = "\033[32m"}
    if ColorChosen == "--color=yellow"{
    color = "\033[33m"}
    if ColorChosen == "--color=blue"{
    color = "\033[34m"}
    if ColorChosen == "--color=purple"{
    color = "\033[35m"}
    if ColorChosen == "--color=cyan"{
    color = "\033[36m"}
    if ColorChosen == "--color=white"{
    color = "\033[37m"}
    if ColorChosen == "--color=orange"{
        color = "\033[91m"}
    if indexofcolorwords == -2 {
        fmt.Println(string(color), charRowData)
    } else {
    fmt.Print(string(color), charRowData)
    }
    color = "\033[0m"
    fmt.Print(string(color))
}