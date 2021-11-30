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
	"math"
	"math/big"
	"strings"

	"github.com/atotto/clipboard"
)

var (
	ascii        = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	unicodeChars = ranges(
		between(384, 535), between(884, 974), between(1025, 1169), between(1328, 1423), between(1488, 1514),
		between(1575, 1610), between(1920, 1969), between(2304, 2431), between(2433, 2554), between(2944, 3071),
		between(3456, 3583), between(3585, 3654), between(3712, 3839), between(3840, 4031), between(4096, 4255),
		between(4256, 4351), between(4608, 4991), between(6016, 6143), between(6144, 6319), between(7936, 8190),
	)
)

func main() {
	out := flag.String("out", "ascii", "specify the output encoding ([ascii, unicode, hex, base32, base64])")
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

func between(start, end int) []int {
	out := []int{}
	for i := start; i <= end; i++ {
		out = append(out, i)
	}
	return out
}

func ranges(rngs ...[]int) []int {
	out := []int{}
	for _, rng := range rngs {
		out = append(out, rng...)
	}

	return out
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
	case "unicode":
		return getUnicode(len(random))
	default:
		log.Fatalf("invalid output type %q", out)
	}

	return ""
}

func getUnicode(length int) string {
	runes := []rune{}
	for i := 0; i < length; i++ {
		i, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		if err != nil {
			log.Fatalf("error generating random unicode character: %v", err)
		}

		runes = append(runes, rune(unicodeChars[int(i.Int64())%len(unicodeChars)]))
	}

	return string(runes)
}

func getASCII(set string, random []byte) string {
	setRunes := []rune(set)
	output := make([]rune, len(random))

	for i := 0; i < len(random); i++ {
		output[i] = setRunes[random[i]%byte(len(setRunes))]
	}

	return string(output)
}
