package main

import (
	"crypto/aes"
	"encoding/hex"
)

const sixteen = "it's should be16"

func encryptAes(token string) (string, error) {
	cip, err := aes.NewCipher([]byte(sixteen))
	if err != nil {
		return "", err
	}
	length := (len(token) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, token)
	pad := byte(len(plain) - len(token))
	for i := len(token); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, cip.BlockSize(); bs <= len(token); bs, be = bs+cip.BlockSize(), be+cip.BlockSize() {
		cip.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return hex.EncodeToString(encrypted), nil
}

func decryptAes(hexStr string) (string, error) {
	encrypted, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	cip, err := aes.NewCipher([]byte(sixteen))
	if err != nil {
		return "", err
	}

	decrypted := make([]byte, len(encrypted))
	for bs, be := 0, cip.BlockSize(); bs < len(encrypted); bs, be = bs+cip.BlockSize(), be+cip.BlockSize() {
		cip.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return string(decrypted[:trim]), nil
}

type user struct {
	Id   string
	Tier string
	Win  int
	Lose int
}

func main() {

}
