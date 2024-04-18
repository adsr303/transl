package transl

import (
	_ "embed"
	"encoding/hex"
	"log"
)

//go:embed ebcdic2unicode.hex
var hexMapping string

var ebcdic2unicode [256]rune

func init() {
	bytes, err := hex.DecodeString(hexMapping)
	if err != nil {
		log.Fatalf("invalid EBCDIC mapping file: %v", err)
	}
	if len(bytes) != len(ebcdic2unicode) {
		log.Fatal("invalid EBCDIC mapping file")
	}
	for i, b := range bytes {
		ebcdic2unicode[i] = rune(b)
	}
}

// DecodeEbcdic converts EBCDIC text in `bytes` to UTF-8 string.
func DecodeEbcdic(bytes []byte) string {
	var result = make([]rune, 0, len(bytes))
	for _, b := range bytes {
		result = append(result, ebcdic2unicode[b])
	}
	return string(result)
}
