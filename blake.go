package main

import (
	"fmt"

	"github.com/dchest/blake2b"
)

func main() {
	var testString string
	fmt.Println("enter the text to be hashed")
	fmt.Scanln(&testString)
	s256 := blake2b.New256()
	s512 := blake2b.New512()
	s256.Write([]byte(testString))
	fmt.Printf("\nSum256 returns a 32-byte BLAKE2b hash of data - %s\n %x\n\n", testString, s256.Sum(nil))
	fmt.Printf("\nSum512 returns a 64-byte BLAKE2b hash of data - %s\n %x\n\n", testString, s512.Sum(nil))
}
