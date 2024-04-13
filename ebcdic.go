package transl

import (
	_ "embed"
	"log"
	"strconv"
)

//go:embed ebcdic2unicode.hex
var hexMapping string

var ebcdic2unicode [256]rune

func init() {
	for i := range 256 {
		n, err := strconv.ParseUint(hexMapping[2*i:2*(i+1)], 16, 8)
		if err != nil {
			log.Fatalf("initializing EBCDIC to Unicode mapping: %v", err)
		}
		ebcdic2unicode[i] = rune(n)
	}
}

// DecodeEbcdic converts EBCDIC text in `bytes“ to UTF-8 string.
func DecodeEbcdic(bytes []byte) string {
	var result = make([]rune, 0, len(bytes))
	for _, b := range bytes {
		result = append(result, ebcdic2unicode[b])
	}
	return string(result)
}
