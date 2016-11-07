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
	args arguments
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

func printRandom(random []byte) {
	switch args.outputType {
	case "hex":
		fmt.Println(hex.EncodeToString(random))
	case "base64":
		fmt.Println(base64.StdEncoding.EncodeToString(random))
	}
}

func parseArgs() {
	args = arguments{}

	flag.StringVar(&args.outputType, "out", "hex", "specify the output encoding ([hex, base64])")
	flag.IntVar(&args.length, "len", 20, "specify the output length")
	flag.Parse()
}
