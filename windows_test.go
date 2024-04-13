package transl

import (
	_ "embed"
	"testing"
)

//go:embed windows_utf16le_bom.txt
var windowsUtf16leWithBom []byte

//go:embed windows_ascii.txt
var windowsAscii []byte

func TestDetectTextEncoding(t *testing.T) {
	tt := []struct {
		text []byte
		enc  Encoding
	}{
		{
			text: windowsUtf16leWithBom,
			enc:  Utf16LE,
		},
		{
			text: windowsAscii,
			enc:  Unknown,
		},
	}
	for _, tc := range tt {
		enc := DetectTextEncoding(tc.text)
		if enc != tc.enc {
			t.Errorf("expected %v, got %v", tc.enc, enc)
		}
	}
}
