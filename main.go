package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	passwordLen = 32 // Limit for replication
)

var (
	charList = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	generate bool
)

func init() {
	flag.BoolVar(&generate, "generate", false, "Generate a new random password")
	flag.Parse()
}

// Hash generates a MySQL compathible hash
func Hash(input string) string {
	in := []byte(input)

	// sha1(sha1(input))
	inner := sha1.Sum(in)
	hash := sha1.Sum(inner[:])

	// MySQL needs an upper-cased hash
	encoded := hex.EncodeToString(hash[:])
	return fmt.Sprintf("*%s", strings.ToUpper(encoded))

}

// NewPass generates a random password
func NewPass() string {
	var pass [passwordLen]byte

	for i := range pass {
		val, err := rand.Int(rand.Reader, big.NewInt(int64(passwordLen)))
		if err != nil {
			panic(err)
		}
		pass[i] = charList[val.Int64()]
	}
	return string(pass[:])
}

func main() {

	flag.Arg(0)

	if !generate {
		fmt.Printf("hash:\t\t%s\n", Hash(os.Args[1]))
	} else {
		pass := NewPass()
		fmt.Printf("password:\t%s\nhash:\t\t%s\n", pass, Hash(pass))
	}

}
