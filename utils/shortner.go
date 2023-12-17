package utils

import (
	b64 "encoding/base64"
	"fmt"
)

func GenerateShortURL(longURL string) string {
    // shortening logic
	encoded_url := b64.StdEncoding.EncodeToString([]byte(longURL))

	fmt.Printf("the original url is: %v and the short url is now %v", longURL, encoded_url)


    // This is a simplistic example using the first 5 characters of the long URL
    return encoded_url
}