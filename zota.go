package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	MerchantID   = "BUGBOUNTY232"
	APISecretKey = "5f4a6fcf-9048-4a0b-afc2-ed92d60fb1bf"
	Currency     = "USD"
	EndpointID   = "402334"
)

var (
	BaseURL = "https://api.zotapay-stage.com"
)

/*const (
	MGClient = require("../zotasdk/client");
	MGDepositRequest = require("../zotasdk/mg_requests/deposit_request")
)*/

type DepositRequest struct {
	MerchantOrderID     string `json:"merchantOrderID"`
	MerchantOrderDesc   string `json:"merchantOrderDesc"`
	OrderAmount         string `json:"orderAmount"`
	OrderCurrency       string `json:"orderCurrency"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerFirstName   string `json:"customerFirstName"`
	CustomerLastName    string `json:"customerLastName"`
	CustomerAddress     string `json:"customerAddress"`
	CustomerCountryCode string `json:"customerCountryCode"`
	CustomerCity        string `json:"customerCity"`
	CustomerZipCode     string `json:"customerZipCode"`
	CustomerPhone       string `json:"customerPhone"`
	CustomerIP          string `json:"customerIP"`
	RedirectURL         string `json:"redirectUrl"`
	CallbackURL         string `json:"callbackUrl"`
	CheckoutURL         string `json:"checkoutUrl"`
	Signature           string `json:"signature"`
}

type DepositResponse struct {
	Code string `json:"code"`
	Data struct {
		DepositURL      string `json:"depositUrl"`
		MerchantOrderID string `json:"merchantOrderID"`
		OrderID         string `json:"orderID"`
	} `json:"data"`
	Message string `json:"message"`
}

func generateSignature(ndpointID string, merchantOrderID string, orderAmount string, customerEmail string, MerchantSecretKey string) string {
	signatureString := fmt.Sprintf("%s%s%s%s%s", ndpointID, merchantOrderID, orderAmount, customerEmail, MerchantSecretKey)
	hash := sha256.Sum256([]byte(signatureString))
	return hex.EncodeToString(hash[:])

	//EndpointID + merchantOrderID + orderAmount + customerEmail + MerchantSecretKey
	//orderId, MerchantID, APISecretKey, EndpointID
	//EndpointID+merchantOrderID+orderAmount+customerEmail+MerchantSecretKey
}

func makeDepositRequest(request *DepositRequest) (*DepositResponse, error) {
	url := fmt.Sprintf("%s/api/v1/deposit/request/%s/", BaseURL, EndpointID)

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 response status: %d, body: %s", resp.StatusCode, string(body))
	}

	var depositResponse DepositResponse
	err = json.Unmarshal(body, &depositResponse)
	if err != nil {
		return nil, fmt.Errorf("faled to unmarshal response: %w", err)
	}

	return &depositResponse, nil
}

func main() {
	orderID := MerchantID
	orderAmount := "500.00"
	orderCurrency := Currency
	customerEmail := "customer@email-address.com"
	request := &DepositRequest{
		MerchantOrderID:     orderID,
		MerchantOrderDesc:   "Test order",
		OrderAmount:         orderAmount,
		OrderCurrency:       orderCurrency,
		CustomerEmail:       customerEmail,
		CustomerFirstName:   "John",
		CustomerLastName:    "Doe",
		CustomerAddress:     "5/5 Moo 5 Thong Nai Pan Noi Beach, Baan Tai, Koh Phangan",
		CustomerCountryCode: "TH",
		CustomerCity:        "Surat Thani",
		CustomerZipCode:     "84280",
		CustomerPhone:       "+66-77999110",
		CustomerIP:          "103.106.8.104",
		RedirectURL:         "https://www.example-merchant.com/payment-return/",
		CallbackURL:         "https://www.example-merchant.com/payment-callback/",
		CheckoutURL:         "https://www.example-merchant.com/account/deposit/?uid=e139b447",
		Signature:           generateSignature(EndpointID, orderID, orderAmount, customerEmail, APISecretKey),
	}

	response, err := makeDepositRequest(request)
	if err != nil {
		fmt.Printf("Failed to make deposit request: %v\n", err)
		return
	}

	code, err := strconv.Atoi(response.Code)
	if err != nil {
		fmt.Printf("Failed to convert response code to integer: %v\n", err)
		return
	}

	if code == 200 {
		fmt.Printf("Deposit URL: %s\n", response.Data.DepositURL)
	} else {
		fmt.Printf("Error: %s\n", response.Message)
	}
}
