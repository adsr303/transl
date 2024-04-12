package transl

import "fmt"

func Example() {
	// "ÅÂÃÄÉÃ@£\u0096@ä\u0095\u0089\u0083\u0096\u0084\u0085"
	input := []byte{197, 194, 195, 196, 201, 195, 64, 163, 150, 64, 228, 149, 137, 131, 150, 132, 133}
	fmt.Println(DecodeEbcdic(input))
	// output: EBCDIC to Unicode
}
