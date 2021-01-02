package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "Zmx1ZmZ5X3Rhbms="
	out, _ := base64.StdEncoding.DecodeString(s)
	fmt.Println(string(out))
}
