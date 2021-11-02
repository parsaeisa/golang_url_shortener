package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func UrlShortener(original_url string) string {
	urlHashBytes := sha256Of(original_url)
	generate_number := new(big.Int).SetBytes(urlHashBytes).Uint64()

	short_url := base58Encoded([]byte(fmt.Sprintf("%d", generate_number)))

	return short_url
}
