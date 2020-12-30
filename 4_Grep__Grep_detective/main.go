package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
)

type level struct {
	explanation string
	hint        string
	solution    string
}

var levels = []level{
	{
		explanation: `
We got a hint that his phone number might be next to the word 'nmbr' in the wordlist.
		`,
		hint: `
You need to use grep to find the line containing the keyword 'nmbr'.
		`,
		solution: `
grep nmbr wordlist
		`,
	},
	{
		explanation: `
There is an important number that you might need later.
The number is the amount of lines in the wordlist that contain the word 'bean' + the last digit of his phone number.
		`,
		hint: `
There is a flag in grep that returns the amount of matches.
`,
		solution: "grep -c \"bean\" wordlist",
	}, {
		explanation: `
Apparently he hid his middle name somewhere in the file called 'middle names.list'.
It's in the line that *ends* with the number from the previous task.
		`,
		hint: `
You can use this character $ somewhere in your search string. You might need to escape it.
		`,
		solution: `
grep "THE_NUMBER\$" "middle names.list"
You might not need to use this: \
		`,
	},
	{
		explanation: `
We heard from a reliable source that his first name contains the letter 'g'.
This might be enough information to find his full name!
		`,
		hint: `
You need to string together multiple things here. If you don't know where to look, find something with full names.
`,
		solution: `
grep "[g].MIDDLE_NAME" "book club attendees.textfile"
	`,
	},
}

func findOutIfPlayerWon() {
	fmt.Print(colorRed)
	fmt.Println("This was all the information we could give you.")
	fmt.Println("Do you have the real full name of Doctor Death?")
	fmt.Print(colorReset + "Full name: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print(colorRed)
	if text == "Flemming Elmer Jensen\n" {
		fmt.Println("Congrats! You did it!")
	} else {
		fmt.Println("That's not right...")
	}
	os.Exit(0)
}

func displayLevels() {
	fmt.Println(colorYellow + "\n(Type a number to pick an option or press ENTER to continue)" + colorReset)
	for i, level := range levels {
		fmt.Printf(colorGreen+"== Level %s =="+colorReset, strconv.Itoa(i+1))
		fmt.Println(level.explanation)
		fmt.Println("1. Hint")
		fmt.Println("2. Solution")
		for {

			text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

			if text == "1\n" {
				fmt.Println(level.hint)
			} else if text == "2\n" {
				fmt.Println(level.solution)
			} else {
				break
			}
		}
	}
}

func main() {
	fmt.Print(colorRed)
	fmt.Println("Welcome detective!")
	fmt.Println("Today we recovered the directory 'Documents' from Doctor Death.\nIt’s your job to find out what his real identity is. ")
	fmt.Print(colorReset)
	displayLevels()
	findOutIfPlayerWon()
}
