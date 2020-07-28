package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"proxypay/model"
	"strconv"
)

var urlProxyPay = flag.String("urlproxypay", "", "url to acess proxypay")
var token = flag.String("token", "", "token to request proxy pay")

//GenerateID This endpoint deletes a reference with given Id.
func GenerateID() (int, error) {

	url := fmt.Sprintf("%s/%s", *urlProxyPay, "reference_ids")
	log.Println(url)
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Add("Authorization", "Token "+*token)
	req.Header.Add("Accept", "application/vnd.proxypay.v2+json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return 0, err
	}

	b, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return 0, errors.New(string(b))
	}

	id, _ := strconv.Atoi(string(b))
	return id, nil
}

//GeneratedRef  create Id in proxypay.
func GeneratedRef(reference model.Reference, id int) error {
	buffer := bytes.Buffer{}
	dec := json.NewEncoder(&buffer)
	dec.Encode(&reference)

	url := fmt.Sprintf("%s/%s/%d", *urlProxyPay, "references", id)
	log.Println(url)
	req, _ := http.NewRequest(http.MethodPut, url, &buffer)
	req.Header.Add("Authorization", "Token "+*token)
	req.Header.Add("Accept", "application/vnd.proxypay.v2+json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return errors.New(string(b))
	}

	return nil
}

//DeleteRef This endpoint deletes a reference with given Id.
func DeleteRef(reference int) error {
	url := fmt.Sprintf("%s/%s/%d", *urlProxyPay, "references", reference)
	log.Println(url)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Add("Authorization", "Token "+*token)
	req.Header.Add("Accept", "application/vnd.proxypay.v2+json")

	resp, err := http.DefaultClient.Do(req) //send request
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(b))
	}

	return nil
}

//GetPayments This endpoint returns any Payment events stored on the server that were not yet Acknowledged by the client application.
func GetPayments(n int) (*[]model.Payment, error) {
	var url string
	if n > 0 {
		url = fmt.Sprintf("%s/%s?n=%d", *urlProxyPay, "payments", n)
	} else {
		url = fmt.Sprintf("%s/%s", *urlProxyPay, "payments")
	}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Token "+*token)
	req.Header.Add("Accept", "application/vnd.proxypay.v2+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return nil, errors.New(string(b))
	}

	payments := []model.Payment{}
	json.Unmarshal(b, &payments)

	return &payments, err
}

//ConfirmPayment  This endpoint is used to acknowledge that a payment was processed
func ConfirmPayment(paymentID int) error {
	url := fmt.Sprintf("%s/%s/%d", *urlProxyPay, "payments", paymentID)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Add("Authorization", "Token "+*token)
	req.Header.Add("Accept", "application/vnd.proxypay.v2+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return errors.New(string(b))
	}

	return nil
}

/*
MockPayment This is only available on the Sandbox environment
and produces a simulated Payment event, as if originated directly from Multicaixa.
*/
func MockPayment(payment model.MockPayment) error {
	buffer := bytes.Buffer{}
	json.NewEncoder(&buffer).Encode(&payment)

	url := fmt.Sprintf("%s/%s", *urlProxyPay, "payments")
	log.Println(url)
	req, _ := http.NewRequest(http.MethodPost, url, &buffer)
	req.Header.Add("Authorization", "Token "+*token)
	req.Header.Add("Accept", "application/vnd.proxypay.v2+json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(b))
	}

	return nil
}
