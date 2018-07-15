package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"bytes"
)

var key32Text = "12345678901234567890123456789012"
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func Encrypt(source,key string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key32Text), err)
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)

	srouceByte := []byte(source)
	ciphertext := make([]byte, len(srouceByte))
	cfb.XORKeyStream(ciphertext, srouceByte)
	//fmt.Printf("%s=>%x\n", source, ciphertext)
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// Decrypt the source string using AES 256 algorithm
func Decrypt(source,key string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key32Text), err)
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)

	srouceByte, _ := base64.URLEncoding.DecodeString(source)
	decryptCopy := make([]byte, len(srouceByte))
	cfbdec.XORKeyStream(decryptCopy, srouceByte)
	//fmt.Printf("%x=>%s\n", source, decryptCopy)
	return string(decryptCopy)
}
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key string) (string) {
	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	pkcs7 := PKCS7Padding([]byte(origData), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keybyte[:blockSize])
	crypted := make([]byte, len(pkcs7))
	blockMode.CryptBlocks(crypted, pkcs7)
	return base64.URLEncoding.EncodeToString(crypted)
}
func AesDecrypt(crypted, key string) (string) {
	cryptedbyte,err := base64.URLEncoding.DecodeString(crypted)
	if err != nil {
		panic(err)
	}
	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keybyte[:blockSize])
	origData := make([]byte, len(cryptedbyte))
	blockMode.CryptBlocks(origData, cryptedbyte)
	origData = PKCS7UnPadding(origData)
	return string(origData)
}
