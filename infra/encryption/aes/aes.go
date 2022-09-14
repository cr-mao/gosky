package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func AesEncrypt(text string, key []byte) (string, error) {
	blockSize := aes.BlockSize //AES的分组大小为16
	src := []byte(text)
	src = ZeroPadding(src, blockSize) //填充
	encrypted := make([]byte, len(src))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCBCEncrypter(block, key) //CBC分组模式加密
	encrypter.CryptBlocks(encrypted, src)
	return hex.EncodeToString(encrypted), nil
}

func AesDecrypt(text string, key []byte) (string, error) {
	src, err := hex.DecodeString(text) //转为[]byte
	if err != nil {
		return "", err
	}
	decrypted := make([]byte, len(src))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	edecrypter := cipher.NewCBCDecrypter(block, key) //CBC分组模式解密
	edecrypter.CryptBlocks(decrypted, src)
	out := ZeroUnPadding(decrypted) //反填充
	return string(out), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding) //用0填充
	return append(ciphertext, padtext...)
}

//ZeroUnPadding 这种方法不严谨，末尾的0不一定全是padding出来的
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0) //截掉尾部连续的0
		})
}

// aes加密, 分组模式ctr
func AesCtrEncrypt(plainText, key []byte) []byte {
	// 1. 建一个底层使用aes的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用ctr分组接口
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	// 4. 加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)
	return cipherText
}

// aes解密
func AesCtrDecrypt(cipherText, key []byte) []byte {
	// 1. 建一个底层使用aes的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用ctr模式解密的接口
	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	// 3. 解密
	message := make([]byte, len(cipherText))
	stream.XORKeyStream(message, cipherText)
	return message
}
