package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"unicode"
)

func IndexHash(input string, index int) string {
	base := Hash(input)
	result := []rune{}
	for _, i := range base {
		if len(result) == index {
			break
		}

		if unicode.IsLetter(i) || unicode.IsNumber(i) {
			result = append(result, i)
		}
	}
	return string(result)
}

func Hash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(hash)
}
