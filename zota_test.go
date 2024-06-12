package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeDepositRequest(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := DepositResponse{
			Code:    "200",
			Message: "Success",
			Data: struct {
				DepositURL      string `json:"depositUrl"`
				MerchantOrderID string `json:"merchantOrderID"`
				OrderID         string `json:"orderID"`
			}{
				DepositURL:      "https://mocked-url.com/deposit",
				MerchantOrderID: "order123",
				OrderID:         "order456",
			},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	// Override the BaseURL to use the mock server
	originalBaseURL := BaseURL
	BaseURL = mockServer.URL
	defer func() { BaseURL = originalBaseURL }()

	// Create a deposit request
	request := &DepositRequest{
		MerchantOrderID:     "order123",
		MerchantOrderDesc:   "Test order",
		OrderAmount:         "100.00",
		OrderCurrency:       "USD",
		CustomerEmail:       "customer@example.com",
		CustomerFirstName:   "John",
		CustomerLastName:    "Doe",
		CustomerAddress:     "123 Street",
		CustomerCountryCode: "US",
		CustomerCity:        "New York",
		CustomerZipCode:     "10001",
		CustomerPhone:       "1234567890",
		CustomerIP:          "127.0.0.1",
		RedirectURL:         "https://example.com/redirect",
		CallbackURL:         "https://example.com/callback",
		CheckoutURL:         "https://example.com/checkout",
		Signature:           generateSignature(EndpointID, "order123", "100.00", "customer@example.com", APISecretKey),
	}

	// Make the deposit request
	response, err := makeDepositRequest(request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate the response
	if response.Code != "200" {
		t.Errorf("Expected response code 200, got %s", response.Code)
	}
	if response.Data.DepositURL != "https://mocked-url.com/deposit" {
		t.Errorf("Expected deposit URL 'https://mocked-url.com/deposit', got %s", response.Data.DepositURL)
	}
	if response.Data.MerchantOrderID != "order123" {
		t.Errorf("Expected merchant order ID 'order123', got %s", response.Data.MerchantOrderID)
	}
	if response.Data.OrderID != "order456" {
		t.Errorf("Expected order ID 'order456', got %s", response.Data.OrderID)
	}
}
