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
}

var (
	args  arguments
	ascii = []rune("1234567890-=qwertyuiop[]asdfghjkl;#zxcvbnm,./¬!£$%^&*()_+QWERTYUIOP{}ASDFGHJKL:@~|ZXCVBNM<>?")
)

func main() {
	parseArgs()

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
	output := make([]rune, len(random))

	for i := 0; i < len(random); i++ {
		output[i] = ascii[random[i]%byte(len(ascii))]
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

func parseArgs() {
	args = arguments{}

	flag.StringVar(&args.outputType, "out", "ascii", "specify the output encoding ([ascii, hex, base32, base64])")
	flag.IntVar(&args.length, "len", 50, "specify the output length")
	flag.Parse()
}
