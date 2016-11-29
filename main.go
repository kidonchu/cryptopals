package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {

	src1 := "1c0111001f010100061a024b53535009181c"
	src2 := "686974207468652062756c6c277320657965"

	dst, err := hexFixedXOR(src1, src2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("dst = %+v\n", dst)
}

func hexToBase64(src string) (string, error) {
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

func hexFixedXOR(src1 string, src2 string) (string, error) {
	decoded1, err := hex.DecodeString(src1)
	if err != nil {
		return "", err
	}
	decoded2, err := hex.DecodeString(src2)
	if err != nil {
		return "", err
	}

	res := make([]byte, len(decoded1))
	for i := 0; i < len(decoded1); i++ {
		res[i] = decoded1[i] ^ decoded2[i]
	}

	encoded := hex.EncodeToString(res)

	return encoded, nil
}
