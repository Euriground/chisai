package utils

import (
	b64 "encoding/base64"
)

func GenerateShortURL(longURL string) string {
    // shortening logic
	encoded_url := b64.StdEncoding.EncodeToString([]byte(longURL))


    // This is a simplistic example using the first 5 characters of the long URL
    return encoded_url
}