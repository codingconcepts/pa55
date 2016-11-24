package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
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

	random := generateRandom()
	printRandom(random)
}

func generateRandom() (random []byte) {
	random = make([]byte, args.length)
	rand.Read(random)

	return
}

func getAscii(random []byte) string {
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
	case "base64":
		fmt.Println(base64.StdEncoding.EncodeToString(random))
	case "ascii":
		fmt.Println(getAscii(random))
	}
}

func parseArgs() {
	args = arguments{}

	flag.StringVar(&args.outputType, "out", "hex", "specify the output encoding ([hex, base64])")
	flag.IntVar(&args.length, "len", 20, "specify the output length")
	flag.Parse()
}
