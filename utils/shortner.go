package utils

import (
	b64 "encoding/base64"
	"fmt"
	"crypto/sha256"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}



func GenerateShortURL(longURL string) string {
    // shortening logic
	urlHashBytes := sha256Of(longURL)
	encoded_url := b64.StdEncoding.EncodeToString([]byte(urlHashBytes))

	fmt.Printf("the original url is: %v and the short url is now %v", longURL, encoded_url)

    // This is a simplistic example using the first 5 characters of the long URL
    return encoded_url[:8]
}