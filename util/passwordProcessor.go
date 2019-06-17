package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// Decrypt : dectypt encoded string
func Decrypt(encrypt string) string {

	encoded, _ := hex.DecodeString(encrypt)
	keyText := "astaxie12798akljzmknm.ahkjkljl;k"
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyText), err)
		os.Exit(-1)
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(encoded))
	cfbdec.XORKeyStream(plaintextCopy, encoded)

	return string(plaintextCopy)
}

// Encrypt : encrypt given string using aes
func Encrypt(plainText string) string {
	textToByte := []byte(plainText)

	keyText := "astaxie12798akljzmknm.ahkjkljl;k"
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyText), err)
		os.Exit(-1)
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherText := make([]byte, len(textToByte))
	cfb.XORKeyStream(cipherText, textToByte)

	hexStr := hex.EncodeToString(cipherText)
	return string(hexStr)
}
