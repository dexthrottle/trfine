package license

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
	"strconv"
	"strings"
	"time"

	"github.com/dexthrottle/trfine/pkg/logging"
)

type LicenseProgram interface {
	CheckLicense()
}

type licenseProgram struct {
	log logging.Logger
}

func NewLicenseProgram(log logging.Logger) LicenseProgram {
	return licenseProgram{
		log: log,
	}
}

const (
	rpineUrlProduction = "https://bot.rpine.xyz:8443/rpine/"
	rpineUrlLocal      = "https://bot.rpine.xyz:8443/rpine_test/"
)

type requestData struct {
	Time    int64  `json:"time"`
	Message string `json:"message"`
}

type responseByBit struct {
	TimeNowByBit string `json:"time_now"`
}

type message struct {
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

type responseMessage struct {
	LicenseAccept string `json:"license_accept"`
	FirstStart    any    `json:"first_start"`
	Time          string `json:"time"`
}

func (l licenseProgram) CheckLicense() {
	msg := message{
		Bot:             "bot-123",
		Version:         "version-123",
		Referal:         "1113",
		Memo:            "memo-123",
		AddressBTC:      "2223",
		AddressBNB_BSC:  "addressBNB_BSC-123",
		AddressUSDT_BSC: "addressUSDT_BSC-123",
		AddressUSDT_TRX: "addressUSDT_TRX-123",
		Time:            time.Now().Unix(),
	}

	publicKey, err := l.bytesToPublicKey([]byte(publicKeyLicense))
	if err != nil {
		l.log.Errorln(err)
	}

	msgEncrypt, err := l.encrypt(&msg, publicKey)
	if err != nil {
		l.log.Errorln(err)
	}

	payload := strings.NewReader(fmt.Sprintf(`{"message": "%s", "time": %d}`, hex.EncodeToString(msgEncrypt), time.Now().Unix()))

	response, err := l.post(rpineUrlLocal, payload)
	if err != nil {
		l.log.Errorf("request error %s'n", err.Error())
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		l.log.Errorf("Error read body request - %s\n", err.Error())
	}

	var reqData requestData
	err = json.Unmarshal(body, &reqData)
	if err != nil {
		l.log.Errorln(err)
	}

	msgToHexDectipt, err := hex.DecodeString(reqData.Message)
	if err != nil {
		l.log.Errorln(err)
	}

	privateKey, err := l.bytesToPrivateKey([]byte(privateKeyLicense))
	if err != nil {
		l.log.Errorln(err)
	}
	decryptByte, err := l.decrypt(msgToHexDectipt, privateKey)
	if err != nil {
		l.log.Errorln(err)
	}

	responseMsg := responseMessage{}
	err = json.Unmarshal(decryptByte, &responseMsg)
	if err != nil {
		l.log.Errorln(err)
	}

	// Получение серверного времени ByBit
	responseBB, err := l.get("https://api-testnet.bybit.com/open-api/v2/public/time")
	if err != nil {
		l.log.Errorln(err)
	}
	defer responseBB.Body.Close()

	bodyBB, err := ioutil.ReadAll(responseBB.Body)
	if err != nil {
		l.log.Errorln("Error read body request - %s", err.Error())
	}

	var resDataBB responseByBit
	err = json.Unmarshal(bodyBB, &resDataBB)
	if err != nil {
		l.log.Errorln(err)
	}

	timeBB, err := strconv.ParseFloat(resDataBB.TimeNowByBit, 64)
	if err != nil {
		l.log.Errorln(err)
	}
	serverBBTime := int64(timeBB * 1000)

	timeServ, err := strconv.ParseFloat(responseMsg.Time, 64)
	if err != nil {
		l.log.Errorln(err)
	}
	serverRPTime := int64(timeServ * 1000)

	licenseAccept, err := strconv.ParseBool(responseMsg.LicenseAccept)
	if err != nil {
		l.log.Errorln(err)
	}

	if serverRPTime >= serverBBTime-3000 && licenseAccept {
		l.log.Println(licenseAccept)
	} else {
		l.log.Println(false)
	}
}

// Отправляет POST запрс на сервер
func (l licenseProgram) post(url string, payload *strings.Reader) (*http.Response, error) {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: transport}
	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l licenseProgram) get(url string) (*http.Response, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: transport}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Возвращает закодированную строку из сообщения
func (l licenseProgram) encrypt(msg *message, key *rsa.PublicKey) ([]byte, error) {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error Marshaling msg to byte - %s", err.Error())
	}
	rng := rand.Reader
	encryptText, err := rsa.EncryptPKCS1v15(rng, key, msgByte)
	if err != nil {
		return nil, err
	}
	return encryptText, nil
}

// Возвращает декодированную строку из сообщения
func (l licenseProgram) decrypt(msgDecrypt []byte, privKey *rsa.PrivateKey) ([]byte, error) {
	rng := rand.Reader
	decryptByte, err := rsa.DecryptPKCS1v15(rng, privKey, msgDecrypt)
	if err != nil {
		return nil, err
	}
	return decryptByte, nil
}

// Преобразует байты PrivateKey в объект PrivateKey
func (l licenseProgram) bytesToPrivateKey(priv []byte) (*rsa.PrivateKey, error) {
	p, _ := pem.Decode(priv)
	key, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// Преобразует байты PublicKey в объект PublicKey
func (l licenseProgram) bytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {
	p, _ := pem.Decode(pub)
	pbKey, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		return nil, err
	}
	key := pbKey.(*rsa.PublicKey)
	return key, nil
}
