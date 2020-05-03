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
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	ascii := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	out := flag.String("out", "ascii", "specify the output encoding ([ascii, hex, base32, base64])")
	length := flag.Int("len", 50, "specify the output length")
	set := flag.String("set", ascii, "character set to use for passwords")
	flag.Parse()

	random, err := generateRandom(*length)
	if err != nil {
		log.Fatalf("error generating random data: %s", err)
	}

	randomStr := prepare(*set, *out, random)
	if err = clipboard.WriteAll(randomStr); err != nil {
		log.Fatalf("error copying password to clipboard: %v", err)
	}

	fmt.Println(strings.Repeat("*", len(randomStr)))
}

func generateRandom(length int) (random []byte, err error) {
	random = make([]byte, length)
	_, err = io.ReadFull(rand.Reader, random)

	return
}

func prepare(set string, out string, random []byte) (output string) {
	switch out {
	case "hex":
		return hex.EncodeToString(random)
	case "base32":
		return base32.StdEncoding.EncodeToString(random)
	case "base64":
		return base64.StdEncoding.EncodeToString(random)
	case "ascii":
		return getASCII(set, random)
	default:
		log.Fatalf("invalid output type %q", out)
	}

	return ""
}

func getASCII(set string, random []byte) string {
	setRunes := []rune(set)
	output := make([]rune, len(random))

	for i := 0; i < len(random); i++ {
		output[i] = setRunes[random[i]%byte(len(setRunes))]
	}

	return string(output)
}
