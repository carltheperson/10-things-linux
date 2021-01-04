package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorReset  = "\033[0m"
)

const minArrowLength = 3
const maxCharsInExplanationBlock = 25

var lsOut []string

var explanations map[string]string = map[string]string{
	"bin":   "/bin contains the executables that retain to very basic system functionality. For example: ls, mkdir and mv.",
	"boot":  "/boot contains the files required to boot up your system.",
	"dev":   "/dev contains special device files for hardware.",
	"etc":   "/etc contains system-wide configuration files.",
	"home":  "/home contains the personal directories for the system users. This is where you find most of the files you use on a day to day basis.",
	"lib":   "/lib contains library files. The kernel modules are also located here.",
	"media": "/media is where external storage automatically is mounted when you plug it in.",
	"mnt":   "/mnt is where storage devices are sometimes mounted. On some distributions they will use the /media for this.",
	"opt":   "/opt contains software that you compile and create yourself.",
	"proc":  "/proc contains information about system resources",
	"root":  "/root is the home directory of the root user.",
	"run":   "/run is used by system processes to store temporary data",
	"sbin":  "/sbin contains binaries for the superuser",
	"srv":   "/srv contains data for servers.",
	"sys":   "/sys contains information about devices connected to your system",
	"tmp":   "/tmp contains temporary files that will be removed on boot",
	"usr":   "/usr contains user libraries, applications and documentation. Here you will also find bin and lib directories.",
	"var":   "/var contains many things including system logs and spools for tasks.",
}

var explanationOrder []string = []string{"bin", "boot", "dev", "etc", "home", "lib", "mnt", "opt", "proc", "root", "run", "sbin", "srv", "sys", "tmp", "usr", "var"}

func init() {
	out, err := exec.Command("ls", "/").Output()
	if err != nil {
		panic(err)
	}

	lsOut = strings.Split(string(out), "\n")
}

func findLongestDirLength() int {
	longestDirLength := 0
	for _, rootDir := range lsOut {
		if len(rootDir) > longestDirLength {
			longestDirLength = len(rootDir)
		}
	}
	return longestDirLength
}

func breakExplanationIntoMultipleLines(explanation string) []string {
	explanationLines := []string{""}
	explanationWords := strings.Split(explanation, " ")

	for len(explanationWords) > 0 {
		latestExplanationLine := &explanationLines[len(explanationLines)-1]
		nextWord := explanationWords[0]
		if len(*latestExplanationLine)+len(nextWord) > maxCharsInExplanationBlock {
			explanationLines = append(explanationLines, nextWord+" ")
		} else {
			*latestExplanationLine += nextWord + " "
		}
		explanationWords = explanationWords[1:]
	}

	return explanationLines
}

func printDirsWithExplanation(explanationLines []string, dirIndex int) {
	longestDirLength := findLongestDirLength()
	currentExplanationLineIndex := 0

	for i, rootDir := range lsOut {
		if i == dirIndex {
			padding := longestDirLength
			extraArrowNeeded := padding - len(rootDir)
			fmt.Print(rootDir)
			arrow := " <" + strings.Repeat("-", minArrowLength+extraArrowNeeded-1)
			fmt.Print(colorYellow + arrow + colorReset + " ")
			if currentExplanationLineIndex < len(explanationLines) {
				fmt.Println(colorBlue + explanationLines[currentExplanationLineIndex] + colorReset)
				explanationLines = explanationLines[1:]
			}

		} else if i > dirIndex {
			padding := longestDirLength + minArrowLength
			spacesNeeded := padding - len(rootDir)
			spaces := strings.Repeat(" ", spacesNeeded)
			fmt.Print(rootDir + spaces)
			if currentExplanationLineIndex < len(explanationLines) {
				fmt.Println("  " + colorBlue + explanationLines[currentExplanationLineIndex] + colorReset)
				explanationLines = explanationLines[1:]
			} else {
				fmt.Println()
			}

		} else {
			fmt.Println(rootDir)
		}
	}

	for currentExplanationLineIndex < len(explanationLines) {
		padding := longestDirLength + minArrowLength
		spaces := strings.Repeat(" ", padding+2)
		fmt.Println(spaces + colorBlue + explanationLines[currentExplanationLineIndex] + colorReset)
		explanationLines = explanationLines[1:]
	}
}

func explainDirs() {
	for _, dirToExplain := range explanationOrder {
		for i, usersDir := range lsOut {
			if dirToExplain == usersDir {
				printDirsWithExplanation(breakExplanationIntoMultipleLines(explanations[dirToExplain]), i)
				bufio.NewReader(os.Stdin).ReadString('\n')
			}
		}
	}
}

func main() {
	explainDirs()
}
