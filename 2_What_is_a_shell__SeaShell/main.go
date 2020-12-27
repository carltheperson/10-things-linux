package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

var shellUser *user.User

func init() {
	shellUser, _ = user.Current()
}

func generateDirPath() string {
	path, _ := os.Getwd()
	pathSplitByHome := strings.Split(path, shellUser.HomeDir)
	if len(pathSplitByHome[0]) == 0 {
		pathSplitByHome[0] = "~"
	}

	return strings.Join(pathSplitByHome, "")
}

func generateStatusLine() string {
	statusLine := colorGreen

	statusLine += shellUser.Name + " "
	statusLine += generateDirPath()
	statusLine += "> "

	return statusLine + colorReset
}

func changeDir(newDir string) error {
	path, _ := os.Getwd()
	if path[0] == '/' {
		return os.Chdir(newDir)
	}
	return os.Chdir(filepath.Join(path, newDir))
}

func interpretInput(input string) string {
	args := strings.Fields(input)

	if args[0] == "cd" && len(args) >= 2 {
		err := changeDir(args[1])
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		return ""
	}

	if args[0] == "exit" {
		os.Exit(0)
	}

	output, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return "\n"
	}
	return (string(output))
}

func main() {
	// ASCII art from: https://asciiart.website/index.php?art=animals/other%20(water) by: jgs @@ shell- ark @@  11/96
	ASCIIArt, _ := base64.StdEncoding.DecodeString("CiAgICAgICAgICAgXy4tLS0uXwogICAgICAgLiciIi4nL3xcYC4iIicuCiAgICAgIDogIC4nIC8gfCBcIGAuICA6CiAgICAgICcuJyAgLyAgfCAgXCAgYC4nCiAgICAgICBgLiAvICAgfCAgIFwgLicKICAgICAgICAgYC0uX198X18uLSc=")
	fmt.Println(colorRed + string(ASCIIArt))
	fmt.Println("\n   SeaShell by carltheperson")
	fmt.Println(colorReset)

	for {
		fmt.Print(generateStatusLine())
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		output := interpretInput(input)
		fmt.Print(output)
	}
}
