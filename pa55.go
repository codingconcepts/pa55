package main

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
)

type arguments struct {
	length     int
	outputType string
	special    string
}

var (
	args    arguments
	ascii   = "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	special = "-=[];#,./¬!£$%^&*()_+{}:@~|<>?"
)

func main() {
	args = arguments{}

	flag.StringVar(&args.outputType, "out", "ascii", "specify the output encoding ([ascii, hex, base32, base64])")
	flag.StringVar(&args.special, "special", special, "special character set to use in ascii passwords")
	flag.IntVar(&args.length, "len", 50, "specify the output length")
	flag.Parse()

	random, err := generateRandom()
	if err != nil {
		log.Fatalf("error generating random data: %s", err)
	}

	printRandom(random)
}

func generateRandom() (random []byte, err error) {
	random = make([]byte, args.length)
	_, err = io.ReadFull(rand.Reader, random)

	return
}

func getASCII(random []byte) string {
	set := []rune(ascii + args.special)
	output := make([]rune, len(random))
	fmt.Println(string(set))

	for i := 0; i < len(random); i++ {
		output[i] = set[random[i]%byte(len(set))]
	}

	return string(output)
}

func printRandom(random []byte) {
	switch args.outputType {
	case "hex":
		fmt.Println(hex.EncodeToString(random))
	case "base32":
		fmt.Println(base32.StdEncoding.EncodeToString(random))
	case "base64":
		fmt.Println(base64.StdEncoding.EncodeToString(random))
	case "ascii":
		fmt.Println(getASCII(random))
	}
}
