package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

var passPhrase = "8dW0if2M40nVP6L2474AsH2homh31K10Nix366i5"

func GetDecryptedNumber(number string) string {

	if len(number) == 0 {
		return ""
	}

	base64DecodedString, err := base64.StdEncoding.DecodeString(number)

	if err != nil {
		fmt.Println("Error decoding base64 string:", err)
		return ""
	}

	encryptedData := Decrypt(base64DecodedString, passPhrase)

	return string(encryptedData)
}

func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(CreateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func CreateHash(key string) string {

	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	encryptedDLNo := "bNFJlOA8kg8+wXHSjAuBXnbkwhP/wNG0imLX1WcTj857IH6cYpCApzyY+sk="
	dl := GetDecryptedNumber(encryptedDLNo)
	fmt.Println(dl)
}
