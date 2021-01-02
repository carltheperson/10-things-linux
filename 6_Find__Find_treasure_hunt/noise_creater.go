package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxDepth = 4

	maxChildDirs     = 4
	maxAmountOfFiles = 5
)

type dir struct {
	amountOfFiles int
	childDirs     []dir
	currentDepth  int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func expandDir(parentDir dir) dir {
	rand.Seed(time.Now().UnixNano())
	amountOfChildDirs := rand.Intn(maxChildDirs-1) + 1

	for i := 0; i < amountOfChildDirs; i++ {
		if parentDir.currentDepth < maxDepth {
			parentDir.childDirs = append(parentDir.childDirs, expandDir(initializeDir(parentDir.currentDepth)))
		}
	}

	return parentDir
}

func initializeDir(currentDepth int) dir {
	rand.Seed(time.Now().UnixNano())
	return dir{
		amountOfFiles: rand.Intn(maxAmountOfFiles-1) + 1,
		currentDepth:  currentDepth + 1,
	}
}

func createFilesFromDir(directory dir, path string) {
	dirName := generateName()
	err := os.Mkdir(path+"/"+dirName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	for i := 0; i < directory.amountOfFiles; i++ {
		file, _ := os.Create(path + "/" + dirName + "/" + generateName() + generateFileEnding())
		file.Write([]byte(generateFileContents()))
		file.Close()
	}

	for _, v := range directory.childDirs {
		createFilesFromDir(v, path+"/"+dirName)
	}

}

func generateName() string {
	adjectives := strings.Split("angry shaky deep sick zippy sticky fluffy frozen unholy honest filthy fighting bonkers harsh frisky greedy crawly insane hideous ungodly abusive drunken hateful idiotic twisted useless", " ")
	nouns := strings.Split("idiot toaster legend tank people band car pistol kitty cat berry messiness dragon mediation toilet deer tailpipe color bug mood freak promises mistake master guilt world sleazeball failure", " ")
	return adjectives[rand.Intn(len(adjectives))] + "_" + nouns[rand.Intn(len(nouns))]
}

func generateFileContents() string {
	loremIpsum := strings.Split("aliquam egestas varius nunc aliquam dictum magna ut porttitor porttitor metus velit tempus sem non placerat massa metus in tortor Etiam bibendum augue vel blandit dictum nulla mauris laoreet leo", " ")

	fileContents := ""

	decider := rand.Intn(3)
	switch decider {
	case 0:
		ipsumLength := rand.Intn(20)
		for i := 0; i < ipsumLength; i++ {
			fileContents += loremIpsum[rand.Intn(len(loremIpsum))] + " "
		}
	case 1:
		fileContents = strconv.Itoa(rand.Intn(1000000000000000000))
	case 2:
		fileContents = generateName() + "_" + generateName()
	}

	return fileContents
}

func generateFileEnding() string {
	fileEndings := []string{".txt", ".text", ".file", ".secret", ".link", ".list", ".txt", ".file", ".zip", ".db", ".doc", "", "", ""}
	return fileEndings[rand.Intn(len(fileEndings))]
}

func main() {
	fmt.Println("Are you sure you want to create noise? This might overwrite what is already there.")
	fmt.Println("y/n")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if text != "y\n" && text != "yes\n" {
		return
	}
	os.Mkdir("noise", os.ModePerm)
	createFilesFromDir(expandDir(initializeDir(0)), "./noise")
}
