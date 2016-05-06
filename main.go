package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func hash(input string) string {
	in := []byte(input)

	// sha1(sha1(input))
	inner := sha1.Sum(in)
	hash := sha1.Sum(inner[:])

	// MySQL needs an upper-cased hash
	encoded := hex.EncodeToString(hash[:])
	return fmt.Sprintf("*%s", strings.ToUpper(encoded))

}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Usage: mysql-hash password")
		os.Exit(1)
	}

	fmt.Println(hash(os.Args[1]))

}
