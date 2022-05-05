package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"

	rsakeys "github.com/dexthrottle/trfine/internal/rsa_keys"
)

const (
	rpineUrl      = "https://bot.rpine.xyz:8443/rpine/"
	rpineUrlLocal = ""
)

type RequestData struct {
	time    int
	message string
}

type Message struct {
	bot             string
	version         string
	referal         string
	memo            string
	addressBTC      string
	addressBNB_BSC  string
	addressUSDT_BSC string
	addressUSDT_TRX string
	time            int
}

func CheckLicense() {
	fmt.Println(rsakeys.PrivateKey)
}

func encrypt(secretMessage string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label)
	if err != nil {
		log.Println(err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func decrypt(cipherText string, privKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Plaintext:", string(plaintext))
	return string(plaintext)
}
