package VavvePaymentsGolang

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/pbkdf2"
)

func AuthorizePayment(messageLayer MessageLayer, apiKey string, apiSecret string) PayemntResponse {

	enc := Encrypting(messageLayer.Account, apiSecret, apiKey)
	messageLayer.Account = enc

	enc = Encrypting(messageLayer.Cvv2, apiSecret, apiKey)
	messageLayer.Cvv2 = enc

	jsonReq, err := json.Marshal(messageLayer)

	req, err := http.NewRequest(http.MethodPost, "http://164.90.146.112:8062/payments/authorize", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-SECRET", apiSecret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var payemntResponse PayemntResponse

	json.Unmarshal(bodyBytes, &payemntResponse)

	return payemntResponse
}

func CapturePaymentPost(capturePayment CapturePayment, apiKey string, apiSecret string) PayemntResponse {

	jsonReq, err := json.Marshal(capturePayment)

	req, err := http.NewRequest(http.MethodPost, "http://164.90.146.112:8062/payments/capture", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-SECRET", apiSecret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var payemntResponse PayemntResponse

	json.Unmarshal(bodyBytes, &payemntResponse)

	return payemntResponse
}

func AuthorizeAndCapturePayment(messageLayer MessageLayer, apiKey string, apiSecret string) PayemntResponse {

	enc := Encrypting(messageLayer.Account, apiSecret, apiKey)
	messageLayer.Account = enc

	enc = Encrypting(messageLayer.Cvv2, apiSecret, apiKey)
	messageLayer.Cvv2 = enc

	jsonReq, err := json.Marshal(messageLayer)

	req, err := http.NewRequest(http.MethodPost, "http://164.90.146.112:8062/payments/authorizeandcapture", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-SECRET", apiSecret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var payemntResponse PayemntResponse

	json.Unmarshal(bodyBytes, &payemntResponse)

	return payemntResponse
}

func RefundAuthorizePayment(refundPayment RefundPayment, apiKey string, apiSecret string) PayemntResponse {

	jsonReq, err := json.Marshal(refundPayment)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://164.90.146.112:8062/payments/authorize/%T/refund", refundPayment.PaymentTransactionId), bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-SECRET", apiSecret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var payemntResponse PayemntResponse

	json.Unmarshal(bodyBytes, &payemntResponse)

	return payemntResponse
}

func RefundCapturePayment(refundPayment RefundPayment, apiKey string, apiSecret string) PayemntResponse {

	jsonReq, err := json.Marshal(refundPayment)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://164.90.146.112:8062/payments/capture/%T/refund", refundPayment.PaymentTransactionId), bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-SECRET", apiSecret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var payemntResponse PayemntResponse

	json.Unmarshal(bodyBytes, &payemntResponse)

	return payemntResponse
}

func VoidAuthorizePayment(voidPayment VoidPayment, apiKey string, apiSecret string) PayemntResponse {

	jsonReq, err := json.Marshal(voidPayment)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://164.90.146.112:8062/payments/authorize/%T/void", voidPayment.PaymentTransactionId), bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-SECRET", apiSecret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var payemntResponse PayemntResponse

	json.Unmarshal(bodyBytes, &payemntResponse)

	return payemntResponse
}

func VoidCapturePayment(voidPayment VoidPayment, apiKey string, apiSecret string) PayemntResponse {

	jsonReq, err := json.Marshal(voidPayment)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://164.90.146.112:8062/payments/capture/%T/void", voidPayment.PaymentTransactionId), bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-SECRET", apiSecret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var payemntResponse PayemntResponse

	json.Unmarshal(bodyBytes, &payemntResponse)

	return payemntResponse
}

func VoidRefundPayment(voidPayment VoidPayment, apiKey string, apiSecret string) PayemntResponse {

	jsonReq, err := json.Marshal(voidPayment)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://164.90.146.112:8062/payments/authorize/%T/voidrefund", voidPayment.PaymentTransactionId), bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-SECRET", apiSecret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var payemntResponse PayemntResponse

	json.Unmarshal(bodyBytes, &payemntResponse)

	return payemntResponse
}

func Encrypting(str string, secret string, salt string) string {
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	spec := pbkdf2.Key([]byte(secret), []byte(salt), 65536, 32, sha256.New)
	block, err := aes.NewCipher([]byte(spec))
	if err != nil {
		fmt.Println(err)
	}
	plainText := PKCS5Padding([]byte(str), block.BlockSize(), len(str))
	cipher := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte, len(plainText))
	cipher.CryptBlocks(cipherText, plainText)
	encrt := base64.StdEncoding.EncodeToString(cipherText)

	return encrt
}

func Decrypting(str string, secret string, salt string) string {
	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	spec := pbkdf2.Key([]byte(secret), []byte(salt), 65536, 32, sha256.New)
	block, err := aes.NewCipher([]byte(spec))
	if err != nil {
		fmt.Println(err)
	}
	cipherTextDecoded, err := base64.StdEncoding.DecodeString(str)
	plainText := []byte(cipherTextDecoded)
	cipher := cipher.NewCBCDecrypter(block, iv)
	cipher.CryptBlocks(plainText, plainText)

	dec := PKCS5UnPadding(plainText)

	return string(dec)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
