echo Try passing as arguments --color red "banana" .
go run main.go --color red "banana"
echo Try passing as arguments --color=red "hello world"
go run main.go --color=red "hello world"
echo Try passing as arguments --color=green "1 + 1 = 2".
go run main.go --color=green "1 + 1 = 2"
echo Try passing as arguments --color=yellow "(%&) ??"
go run main.go --color=yellow "(%&) ??"
echo Try specifying a set of letters to be colored the second until the last letter.
go run main.go --color=purple ontfailme "Dontfailme"
echo Try specifying letter to be colored the second letter.
go run main.go --color=cyan e "hello world"
echo Try specifying a set of letters to be colored just two letters
go run main.go --color=red OM "COME ON!!"
echo Try passing as arguments --color=orange GuYs "HeY GuYs", in order to color GuYs.
go run main.go --color=orange GuYs "HeY GuYs"
echo Try passing as arguments --color=blue B 'RGB()', in order to color just the B.
go run main.go --color=blue B 'RGB()'
echo Try passing as arguments a random string with lower and upper case letters, and a random color in the color flag "--color=".
go run main.go --color=green "SerIOuSLY!!"
echo Try passing as arguments a random string with lower case letters, numbers and spaces, and a random color in the color flag "--color=".
go run main.go --color=black "werwer$##$@ddd"
echo Try passing as arguments a random string with special characters, and a random color in the color flag "--color=", specifying one letter to be colored.
go run main.go --color=red a "random"
echo Try passing as arguments a random string with lower case letters, upper case letters, spaces and numbers with a random color in the color flag "--color=", specifying a set of letters to be colored.
go run main.go --color=red nal "Last%%Question**Finally"
