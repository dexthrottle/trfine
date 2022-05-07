package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	rsakeys "github.com/dexthrottle/trfine/internal/rsa_keys"
)

const (
	rpineUrlProduction = "https://bot.rpine.xyz:8443/rpine/"
	rpineUrlLocal      = "https://bot.rpine.xyz:8443/rpine_test/"
)

type RequestData struct {
	Time    int64  `json:"time"`
	Message string `json:"message"`
}

type Message struct {
	Bot             string `json:"bot"`
	Version         string `json:"version"`
	Referal         string `json:"referal"`
	Memo            string `json:"memo"`
	AddressBTC      string `json:"address_btc"`
	AddressBNB_BSC  string `json:"address_bnb_BSC"`
	AddressUSDT_BSC string `json:"address_USDT_BSC"`
	AddressUSDT_TRX string `json:"address_USDT_TRX"`
	Time            int64  `json:"time"`
}

type ResponseMessage struct {
	LicenseAccept bool  `json:"license_accept"`
	FirstStart    bool  `json:"first_start"`
	Time          int64 `json:"time"`
}

func CheckLicense() {

	msg := Message{
		Bot:             "bot-123",
		Version:         "version-123",
		Referal:         "referal-123",
		Memo:            "memo-123",
		AddressBTC:      "addressBTC-123",
		AddressBNB_BSC:  "addressBNB_BSC-123",
		AddressUSDT_BSC: "addressUSDT_BSC-123",
		AddressUSDT_TRX: "addressUSDT_TRX-123",
		Time:            time.Now().Unix(),
	}
	pubKey := BytesToPublicKey([]byte(rsakeys.PublicKey))

	msgEncrypt := encrypt(&msg, pubKey)
	p := fmt.Sprintf(`{"message": "%s", "time": %d}`, hex.EncodeToString(msgEncrypt), time.Now().Unix())
	fmt.Println(p)
	payload := strings.NewReader(p)

	response, err := post(rpineUrlLocal, payload)
	if err != nil {
		log.Printf("request error %s", err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error read body request - %s", err.Error())
	}

	var rData RequestData
	err = json.Unmarshal(body, &rData)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v", rData)

	decryptByte := decrypt(rData.Message, BytesToPrivateKey([]byte(rsakeys.PrivateKey)))
	fmt.Println(string(decryptByte))

}

func post(url string, payload *strings.Reader) (*http.Response, error) {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{Transport: transport}
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		log.Println("Error creating HTTP request")
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	fmt.Printf("%+v\n\n\n", req)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error retrieving tasks from the backend")
		return nil, err
	}

	return resp, nil
}

// Возвращает закодированную строку из json
func encrypt(msg *Message, key *rsa.PublicKey) []byte {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error Marshaling msg to byte - %s", err.Error())
	}
	// label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptPKCS1v15(rng, key, msgByte)
	if err != nil {
		log.Println(err)
	}
	return ciphertext
}

func decrypt(msgDecrypt string, privKey *rsa.PrivateKey) []byte {
	rng := rand.Reader
	decryptByte, err := rsa.DecryptPKCS1v15(rng, privKey, []byte(msgDecrypt))
	if err != nil {
		log.Println(err)
	}
	return decryptByte
}

func BytesToPrivateKey(priv []byte) *rsa.PrivateKey {
	p, _ := pem.Decode(priv)
	key, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	return key
}

func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	p, _ := pem.Decode(pub)
	pbKey, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	key := pbKey.(*rsa.PublicKey)
	return key
}
