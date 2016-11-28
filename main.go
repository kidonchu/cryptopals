package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	msg := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	dst, err := HexToBase64(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("dst = %+v\n", dst)
}

// HexToBase64 converts hexadecimal encoding to base64 encoding
func HexToBase64(src string) (string, error) {
	decoded, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	encoded := base64.StdEncoding.EncodeToString(decoded)
	if err != nil {
		return "", err
	}
	return encoded, nil
}
