package transl

import "slices"

type Encoding int

const (
	Unknown Encoding = iota
	Utf8
	Utf8WithBOM
	Utf16LE
)

var utf16leBom = []byte{0xff, 0xfe}
var utf8Bom = []byte{0xef, 0xbb, 0xbf}

// DetectTextEncoding attempts to guess encoding of text stored in `bytes`.
func DetectTextEncoding(bytes []byte) Encoding {
	if len(bytes) < 2 {
		return Unknown // TODO try to guess - maybe ASCII or UTF-8, probably something plain
	}
	if slices.Equal(bytes[:2], utf16leBom) {
		return Utf16LE
	}
	if len(bytes) < 3 {
		return Unknown // TODO probably short text, again try to guess
	}
	if slices.Equal(bytes[:3], utf8Bom) {
		return Utf8WithBOM
	}
	return Unknown
}
