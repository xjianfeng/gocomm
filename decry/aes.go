package decry

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	log "github.com/xjianfeng/gocomm/logger"
)

var key = []byte("game@123jiamikey")

func DoAesEncrypt(entryStr []byte) string {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	result, err := AesEncrypt(entryStr, key)
	if err != nil {
		log.LogError("DoAesEncrypt Error %s", err.Error())
		return ""
	}
	res := base64.StdEncoding.EncodeToString(result)
	return res
}

func DoAesDecrypt(result string) string {
	encryData, err := base64.StdEncoding.DecodeString(result)
	if err != nil {
		log.LogError("DoAesDecrypt Base64 DecodeString Error %s", err.Error())
		return ""
	}
	origData, err := AesDecrypt(encryData, key)
	if err != nil {
		log.LogError("DoAesDecrypt Error %s", err.Error())
		return ""
	}
	return string(origData)
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
