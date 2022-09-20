package shortener

// Url shortener
// reference from: https://github.com/eddywm/go-shortener-wm

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/google/uuid"
	"github.com/itchyny/base58-go"
)

func GenerateCode(uuid uuid.UUID) string {
	urlHashBytes := sha256Of(uuid.String())
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
